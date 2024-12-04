package interfaces

import "math/big"

type PaymentProcessor interface {
	PayNetworkTax(nodes *[]string, tax *big.Int) bool
	CalculatePayment(size int) *big.Int
}
