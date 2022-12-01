package main

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	senderBlockChainAddress    string
	recipientBlockChainAddress string
	message                    string
}

func (t *Transaction) Print() {
	fmt.Println("sender address:		", t.senderBlockChainAddress)
	fmt.Println("recipient address:	", t.recipientBlockChainAddress)
	fmt.Println("value:			", t.message)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string `json:"senderBlockChainAddress"`
		Recipient string `json:"recipientBlockChainAddress"`
		Message   string `json:"message"`
	}{
		Sender:    t.senderBlockChainAddress,
		Recipient: t.recipientBlockChainAddress,
		Message:   t.message,
	})
}

func NewTransaction(sender string, recipient string, message string) *Transaction {
	return &Transaction{
		senderBlockChainAddress:    sender,
		recipientBlockChainAddress: recipient,
		message:                    message,
	}
}
