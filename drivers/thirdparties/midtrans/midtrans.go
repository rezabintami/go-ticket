package midtrans

import (
	"context"
	_config "ticketing/app/config"
	"ticketing/business/payments"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type TransactionMidtrans struct {
	snapClient snap.Client
}

func NewTransactionMidtrans() payments.Repository {
	return &TransactionMidtrans{
		snapClient: snap.Client{},
	}
}

func (tm *TransactionMidtrans) Transactions(ctx context.Context, transactionDomain *payments.Domain) (payments.DomainResponse, error) {
	tm.snapClient.New(_config.GetConfig().MIDTRANS_SERVER_KEY, midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: "YOUR-ORDER-ID-12345", GrossAmt: 100000},
		CreditCard: &snap.CreditCardDetails{Secure: true},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "John",
			LName: "Doe",
			Email: "john@doe.com",
			Phone: "081234567890",
		},
	}

	snapResp, _ := tm.snapClient.CreateTransaction(req)
	data := fromDomain(*snapResp)
	respDomain := toDomain(*data)
	return respDomain, nil
}
