package bradesco

import (
	"go-brasil-banco/dominio/comum"
	"strconv"
)

// cálculo  do  dígito
func CalculoDigitoNossoNumero(carteira string, nossonumero string) (string, error) {

	strValor := comum.Zeros(carteira, 3) + comum.Zeros(nossonumero, 11)
	data, err := comum.Explode(strValor)

	if err != nil {
		return "", err
	}

	var soma int
	var base = 7
	var peso = 0
	var i = 0
	var dv = ""
	var constMod = 11
	for i, peso = len(data)-1, 2; i >= 0; i, peso = i-1, peso+1 {

		if peso > base {
			peso = 2
		}
		soma += data[i] * peso
	}
	mod11 := soma % constMod
	if mod11 == 1 {
		dv = "P"
	} else if mod11 == 0 {
		dv = "0"
	} else {
		dv = strconv.Itoa(11 - mod11)
	}
	return dv, nil
}
