package models

import "math/big"

type PaymentRequest struct {
	PaymentID string
	Amount    *big.Int
}
