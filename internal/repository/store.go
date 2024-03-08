package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Store struct {
	Querier
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Querier: NewRepository(db),
	}
}

func NewRepository(db *sql.DB) Querier {
	return New(db)
}

// create a transaction with context then call back function to execute queries
func (st *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	que := New(tx)
	err = fn(que)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v \n rollback error : %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (st *Store) LikePostWithTx(ctx context.Context, accountId, postId uuid.UUID) error {
	return st.execTx(ctx, func(q *Queries) error {
		err := q.CreateLikedPost(ctx, CreateLikedPostParams{
			AccountID: accountId,
			PostID:    postId,
		})
		if err != nil {
			return err
		}
		id, err := q.UpdateLikeCountIncrement(ctx, postId)
		if err != nil {
			return err
		}
		err = q.UpdateGetLikeUserDetail(ctx, id)
		if err != nil {
			return err
		}
		err = q.UpdateGiveLikeUserDetail(ctx, accountId)
		if err != nil {
			return err
		}
		return nil
	})
}

func (st *Store) UnlikePostWithTx(ctx context.Context, accountId, postId uuid.UUID) error {
	return st.execTx(ctx, func(q *Queries) error {
		err := q.DeleteLikedPost(ctx, DeleteLikedPostParams{
			AccountID: accountId,
			PostID:    postId,
		})
		if err != nil {
			return err
		}
		id, err := q.UpdateLikeCountDecrement(ctx, postId)
		if err != nil {
			return err
		}
		err = q.UpdateDecreaseGetLikeUserDetail(ctx, id)
		if err != nil {
			return err
		}
		err = q.UpdateDecreaseGiveLikeUserDetail(ctx, accountId)
		if err != nil {
			return err
		}
		return nil
	})
}

func (st *Store) DeletePostWithTx(ctx context.Context, postId uuid.UUID) error {
	return st.execTx(ctx, func(q *Queries) error {
		owener_id, err := q.QueryUserIdfromPost(ctx, postId)
		if err != nil {
			return err
		}
		account_id, err := q.QueryGetAccoutFromLikedByPostId(ctx, postId)
		if err != nil {
			return err
		}
		for _, i := range account_id {
			q.UpdateDecreaseGetLikeUserDetail(ctx, owener_id)
			q.UpdateDecreaseGiveLikeUserDetail(ctx, i)
		}
		err = q.DeleteLikedByPostId(ctx, postId)
		if err != nil {
			return err
		}
		err = q.DeleteImageByPostId(ctx, postId)
		if err != nil {
			return err
		}
		err = q.DeleteCommentByPostId(ctx, postId)
		if err != nil {
			return err
		}
		err = q.DeletePostById(ctx, postId)
		if err != nil {
			return err
		}

		return nil
	})
}

// note: we cah use select for update no key , so pgsql know we would not update
// the primary key so it will no deadlock because we not update the primary key
// constraint to other table and it will be consistent

// intinya kalau kalau ada forign key dan kita insert di table yang puunya
// foregin key contraint , dan kita select for update in transaksi yang lain
// maka akan terjadi deadlock , sehingga kita harus pastikan bahawa select for
// update tidak mengupdate primary key
