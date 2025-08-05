package main

import (
	"strconv"
)

func IsValidaCnpj(cnpj string) bool {
	cnpj = removeNonDigits(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	// Rejeita CNPJ com dígitos iguais
	for i := 1; i < 14; i++ {
		if cnpj[i] != cnpj[0] {
			break
		}
		if i == 13 {
			return false
		}
	}

	base := cnpj[:12]
	d1 := calculateDigit(base, []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	d2 := calculateDigit(base+d1, []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})

	return cnpj == base+d1+d2
}

func calculateDigit(cnpj string, weigths []int) string {
	sum := 0
	for i := 0; i < len(weigths); i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * weigths[i]
	}
	result := sum % 11
	if result < 2 {
		return "0"
	}
	return strconv.Itoa(11 - result)
}
