package midtrans

import (
	"context"
	_config "ticketing/app/config"
	"ticketing/business/payments"

	"github.com/veritrans/go-midtrans"
)

type TransactionMidtrans struct {
	midClient midtrans.Client
}

func NewTransactionMidtrans() payments.Repository {
	return &TransactionMidtrans{
		midClient: midtrans.Client{},
	}
}

func (tm *TransactionMidtrans) Transactions(ctx context.Context, transactionDomain *payments.Domain) (payments.DomainResponse, error) {
	tm.midClient.ServerKey = _config.GetConfig().MIDTRANS_SERVER_KEY
	tm.midClient.ClientKey = _config.GetConfig().MIDTRANS_CLIENT_KEY
	tm.midClient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: tm.midClient,
	}

	req := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionDomain.OrderID,
			GrossAmt: int64(transactionDomain.Amount),
		},
		CreditCard: &midtrans.CreditCardDetail{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: transactionDomain.Name,
			Email: transactionDomain.Email,
		},
	}

	snapTokenResponse, err := snapGateway.GetToken(req)
	if err != nil {
		return payments.DomainResponse{}, err
	}
	data := fromDomain(snapTokenResponse)
	respDomain := toDomain(*data)
	return respDomain, nil
}
