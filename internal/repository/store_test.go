package repository

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/unedtamps/go-backend/util"
)

func createRandomPayment(user CreateUserRow, t *testing.T) Payment{
  args := CreatePaymentParams{
    ID:     util.RandomUUID(),
    UserID: sql.NullString{
      String: user.ID,
      Valid: true,
    },
    PremiumTypeID:util.RandomTypePremium(),
    Amount: 50000,
  }
  payment, err := testQueries.CreatePayment(context.Background(), args)
  require.NoError(t, err)
  require.NotEmpty(t, payment)
  require.Equal(t, args.ID, payment.ID)
  require.Equal(t, args.UserID, payment.UserID)
  require.Equal(t, args.Amount, payment.Amount)

  return *payment
}


func TestUpdateStatusPaymentWithTx(t *testing.T) {
  store := NewStore(testDB)
  n := 10
  arr_user := make([]CreateUserRow, n)
  for i := 0; i < n; i++ {
    arr_user[i] = createRandomUser(t)
  }

  errChan := make(chan error)
  paymentChan := make(chan *Payment)

  for _, v := range arr_user {
    go func (v CreateUserRow)  {
    payment_user_1 := createRandomPayment(v, t)
    paymentRes, err := store.UpdateStatusPaymentWithTx(context.Background(), payment_user_1.ID, payment_user_1.PremiumTypeID, PaymentStatusSuccess)
    errChan <- err
    if err != nil{
        paymentRes, err:= store.UpdatePaymentStatus(context.Background(), UpdatePaymentStatusParams{
          Status: NullPaymentStatus{
            PaymentStatus: PaymentStatusFailed,
            Valid: true,
          },
          ID: payment_user_1.ID,
        })
        if err != nil {
          log.Println(err.Error())

        }
        paymentChan <- paymentRes
      }else {
        paymentChan <- paymentRes
      }

    }(v)
  }

  for i := 0; i < n; i++ {
    err := <-errChan
    paymentRes := <- paymentChan
    if err != nil {
      log.Printf("got this error %v", err)
      require.Equal(t, paymentRes.Status.PaymentStatus, PaymentStatusFailed)
      continue
    }

      require.Equal(t, paymentRes.Status.PaymentStatus, PaymentStatusSuccess)
  }
}
