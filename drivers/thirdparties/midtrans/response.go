package midtrans

import (
	"ticketing/business/payments"

	"github.com/midtrans/midtrans-go/snap"
)

type Response struct {
	Token         string   `json:"token"`
	RedirectURL   string   `json:"redirect_url"`
	StatusCode    string   `json:"status_code,omitempty"`
	ErrorMessages []string `json:"error_messages,omitempty"`
}

func toDomain(resp Response) payments.DomainResponse {
	return payments.DomainResponse{
		Token:         resp.Token,
		RedirectURL:   resp.RedirectURL,
		StatusCode:    resp.StatusCode,
		ErrorMessages: resp.ErrorMessages,
	}
}

func fromDomain(snapResp snap.Response) *Response {
	return &Response{
		Token:         snapResp.Token,
		RedirectURL:   snapResp.RedirectURL,
		StatusCode:    snapResp.StatusCode,
		ErrorMessages: snapResp.ErrorMessages,
	}
}
