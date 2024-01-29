package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func NewRepository(conn *sql.DB) *Store {
	return NewStore(conn)
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

// note: we cah use select for update no key , so pgsql know we would not update
// the primary key so it will no deadlock because we not update the primary key
// constraint to other table and it will be consistent

// intinya kalau kalau ada forign key dan kita insert di table yang puunya
// foregin key contraint , dan kita select for update in transaksi yang lain
// maka akan terjadi deadlock , sehingga kita harus pastikan bahawa select for
// update tidak mengupdate primary key

func (st *Store) UpdateStatusPaymentWithTx(
	ctx context.Context,
	paymentID string,
	premiumTypeID string,
	status PaymentStatus,
) (*Payment, error) {
	var payment *Payment
	err := st.execTx(ctx, func(que *Queries) error {
		var err error

		quota, err := que.GetQuotaByPremiumTypeIdForUpdate(ctx, sql.NullString{
			String: premiumTypeID,
			Valid:  true,
		})

		if err != nil {
			return err
		}

		err = que.DecreaseOneQuota(ctx, quota.ID)
		if err != nil {
			return err
		}

		payment, err = que.UpdatePaymentStatus(ctx, UpdatePaymentStatusParams{
			ID:     paymentID,
			Status: status,
		})
		if err != nil {
			return err
		}
		err = que.ChangeUserType(ctx, ChangeUserTypeParams{
			AcountType: true,
			ID:         payment.UserID,
		})

		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return payment, nil
}
