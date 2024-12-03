package interfaces

import "math/big"

type PaymentProcessor interface {
	PayNetworkTax(nodes *[][32]byte, tax *big.Int) bool
	CalculatePayment(size int) *big.Int
}
