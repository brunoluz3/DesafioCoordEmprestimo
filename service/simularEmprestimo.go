package service

import (
	"math/big"
)

// MonthlyPaymentFloat calcula a parcela (PRICE) usando big.Float.
// pv   = valor presente
// rate = taxa mensal (ex: 0.02)
// n    = número de parcelas

func MonthlyPaymentFloat(pv, rate *big.Float, n int) *big.Float {
	precision := uint(128)
	zero := new(big.Float).SetPrec(precision).SetFloat64(0)
	one := new(big.Float).SetPrec(precision).SetFloat64(1)

	// Se taxa = 0 → valor / n
	if rate.Cmp(zero) == 0 {
		nFloat := new(big.Float).SetPrec(precision).SetFloat64(float64(n))
		return new(big.Float).Quo(pv, nFloat)
	}

	// (1 + rate)
	onePlusRate := new(big.Float).Add(one, rate)

	// (1 + rate)^n
	pow := new(big.Float).SetPrec(precision).Copy(one)
	for i := 0; i < n; i++ {
		pow.Mul(pow, onePlusRate)
	}

	// PRICE:
	// parcela = pv * rate * (1+rate)^n / ((1+rate)^n - 1)
	num := new(big.Float).Mul(pv, rate)
	num.Mul(num, pow)

	den := new(big.Float).Sub(pow, one)

	return new(big.Float).Quo(num, den)
}
