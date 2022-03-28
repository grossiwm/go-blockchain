package transaction

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderAddress    string
	recipientAddress string
	value            float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (transaction *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Printf(" senderAddress		%s\n", transaction.senderAddress)
	fmt.Printf(" recipientAddress	%s\n", transaction.recipientAddress)
	fmt.Printf(" value	%.4f\n", transaction.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderAdress     string  `json:"senderAdress"`
		RecipientAddress string  `json:"recipientAddress"`
		Value            float32 `json:"value"`
	}{
		SenderAdress:     t.senderAddress,
		RecipientAddress: t.recipientAddress,
		Value:            t.value,
	})
}
