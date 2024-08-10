package utils

import (
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentPayload struct {
	OrderId int
	Amount int
	FName string
	Email string
}

func CreatePayment(payload PaymentPayload) (string, error) {
	serverKey := MustGetEnv("MIDTRANS_SERVER_KEY")

	s := snap.Client{}
	s.New(serverKey, midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(payload.OrderId),
			GrossAmt: int64(payload.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: payload.FName,
			Email: payload.Email,
		},
	}

	snapRes, err := s.CreateTransaction(req)
	if err != nil {
		return "", nil
	}
	return snapRes.RedirectURL, nil
}