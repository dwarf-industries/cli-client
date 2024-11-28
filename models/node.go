package models

import "math/big"

type Node struct {
	Name       string
	Ip         string
	Port       string
	Reputation big.Int
}
