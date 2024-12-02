package converters

import "math/big"

func WeiToEth(wei *big.Int) *big.Float {
	weiToEthFactor := new(big.Float).SetFloat64(1e18)
	weiFloat := new(big.Float).SetInt(wei)
	eth := new(big.Float).Quo(weiFloat, weiToEthFactor)

	return eth
}
