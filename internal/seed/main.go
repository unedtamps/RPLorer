package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/util"

	_ "github.com/lib/pq"
)

var querier repository.Queries

type DBENV struct {
	DBDriver         string `mapstructure:"DB_DRIVER"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PosgresHost      string `mapstructure:"POSTGRES_HOST"`
}

func getDbEnv() *DBENV {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		util.Log.Fatal(err)
	}
	env := DBENV{}
	err = viper.Unmarshal(&env)
	if err != nil {
		util.Log.Fatal(err)
	}
	return &env
}

func main() {
	env := getDbEnv()
	db_url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		env.PosgresHost,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDB,
	)
	db, err := sql.Open(env.DBDriver, db_url)
	if err != nil {
		util.Log.Fatal(err)
	}
	querier = *repository.New(db)
	ctx := context.Background()

	if len(os.Args) > 1 && os.Args[1] == "admin" {
		password, _ := util.GenereateHasedPassword("adminrpl12")
		admin, err := querier.CreateAdminAccount(ctx, repository.CreateAdminAccountParams{
			FirstName: "admin",
			LastName:  "unedo",
			Username:  "admin_unedo",
			Email:     "admin@gmail.com",
			Password:  password,
			Role:      repository.RoleADMIN,
			Status:    repository.AccountStatusACTIVE,
		})
		if err != nil {
			util.Log.Info(err)
			return
		}
		err = querier.CreateAccountDetail(ctx, repository.CreateAccountDetailParams{
			AccountID:      admin,
			GiveLike:       0,
			GetLike:        0,
			FollowersCount: 0,
			FollowingCount: 0,
		})
		if err != nil {
			util.Log.Info(err)
			return
		}
		util.Log.Info("success seeding admin account")
		return
	}

	for i := 0; i < 5; i++ {
		// create account
		id, err := seedAccount(ctx)
		if err != nil {
			util.Log.Fatal(err)
		}
		err = seedAccountDetail(ctx, id)
		if err != nil {
			util.Log.Fatal(err)
		}
		fmt.Println("Account ID: ", id)

		// create post for this account
		for j := 0; j < 10; j++ {
			likes := util.RandomInt8()
			comment := util.RandomInt8()
			post_id, err := seedPost(ctx, id, int64(likes), int64(comment))
			if err != nil {
				util.Log.Fatal(err)
			}
			fmt.Println("post_id: ", post_id)

			// create liked post and comment for this post
			for k := 0; k < likes; k++ {
				acc_id, err := seedAccount(ctx)
				if err != nil {
					util.Log.Fatal(err)
				}
				seedAccountDetail(ctx, acc_id)
				err = querier.CreateLikedPost(ctx, repository.CreateLikedPostParams{
					AccountID: acc_id,
					PostID:    post_id,
				})
				querier.UpdateGetLikeUserDetail(ctx, id)
				querier.UpdateGiveLikeUserDetail(ctx, acc_id)

				// add followers
				err = querier.CreateAccountFollow(ctx, repository.CreateAccountFollowParams{
					AccountFollowed:  id,
					AccountFollowing: acc_id,
				})
				if err != nil {
					util.Log.Fatal(err)
				}
				querier.UpdateFollowersCount(ctx, id)
				querier.UpdateFollowingCount(ctx, acc_id)
			}

			// create comment for this post
			new_comment := 0
			for k := 0; k < comment; k++ {
				new_comment++
				id, err := seedAccount(ctx)
				if err != nil {
					util.Log.Fatal(err)
				}
				seedAccountDetail(ctx, id)
				c_id, err := seedComment(ctx, id, post_id, uuid.NullUUID{Valid: false})
				for l := 0; l < 2; l++ {
					new_comment++
					new_id, err := seedAccount(ctx)
					if err != nil {
						util.Log.Fatal(err)
					}
					seedAccountDetail(ctx, new_id)
					seedComment(ctx, new_id, post_id, uuid.NullUUID{Valid: true, UUID: c_id})
				}
			}
			err = querier.UpdatePostDetailCount(ctx, repository.UpdatePostDetailCountParams{
				LikesCount:   int64(likes),
				CommentCount: int64(new_comment),
				ID:           post_id,
			})

		}
		util.Log.Info("Success seeding account: ", id)
	}
}

func seedAccount(ctx context.Context) (uuid.UUID, error) {
	acc, err := querier.CreateAccount(ctx, repository.CreateAccountParams{
		FirstName: util.RandomString(),
		LastName:  util.RandomString(),
		Username:  util.RandomString(),
		Email:     util.RandomEmail(),
		Password:  util.RandomHashedPassword(),
	})
	return acc.ID, err
}

func seedAccountDetail(ctx context.Context, id uuid.UUID) error {
	err := querier.CreateAccountDetail(ctx, repository.CreateAccountDetailParams{
		AccountID:      id,
		GetLike:        0,
		GiveLike:       0,
		FollowersCount: 0,
		FollowingCount: 0,
	})
	return err
}

func seedPost(ctx context.Context, account_id uuid.UUID, like, comment int64) (uuid.UUID, error) {
	id, err := querier.CreatePost(ctx, repository.CreatePostParams{
		AccountID:    account_id,
		Caption:      util.RandomString(),
		LikesCount:   like,
		CommentCount: comment,
	})
	return id, err
}

func seedComment(
	ctx context.Context,
	acc_id, post_id uuid.UUID,
	parrent uuid.NullUUID,
) (uuid.UUID, error) {
	id, err := querier.CreateComment(ctx, repository.CreateCommentParams{
		AccountID: acc_id,
		PostID:    post_id,
		Body:      util.RandomString(),
		ParrentID: parrent,
	})
	return id, err
}
