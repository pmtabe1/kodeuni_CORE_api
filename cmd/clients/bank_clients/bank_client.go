package bank_clients

import "github.com/paulmsegeya/pos/cmd/clients/bank_clients/tcbbank_client"

 
type IBankClienr interface {
}

type BankClient struct {
	TCBBankClient *tcbbank_client.TCBBankClient
}

func New() *BankClient {

	return &BankClient{
		TCBBankClient: tcbbank_client.New(),
	}
}
