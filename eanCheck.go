package main

import "strconv"

func isCorrectEan(possibleEan string) bool {

	// test longueur

	var lenEan int = len(possibleEan)

	if lenEan == 13 || lenEan == 8 {
		// test si uniquement nombre

		var _, errIsNumber = strconv.ParseUint(possibleEan, 10, 64)

		if errIsNumber == nil {

			// test digit check
			var eanDigitless = ""
			var possibleDigit = ""
			switch lenEan {
			case 13:
				eanDigitless = possibleEan[0:12]
				possibleDigit = string(possibleEan[12])
			case 8:
				eanDigitless = possibleEan[0:7]
				possibleDigit = string(possibleEan[7])

			}

			/*if lenEan == 13 {
				eanDigitless = possibleEan[0:12]
				possibleDigit = string(possibleEan[12])
			}
			if lenEan == 8 {
				eanDigitless = possibleEan[0:7]
				possibleDigit = string(possibleEan[7])
			}*/

			var trueDigit = calculateDigitCheck(eanDigitless)
			if trueDigit == possibleDigit {
				return true
			}

			return false
		}
		return false

	}
	return false

}

func calculateDigitCheck(eanDigitCheckless string) string {
	var lenstrCalcul int = len(eanDigitCheckless)
	var factor uint64 = 3
	var somme uint64 = 0

	var digitCheck = ""

	for index := lenstrCalcul - 1; index > -1; index-- {
		var tmpCalcul, _ = strconv.ParseUint(string(eanDigitCheckless[index]), 10, 64)
		somme += tmpCalcul * factor
		factor = 4 - factor
	}

	digitCheck = strconv.FormatUint((10-(somme%10))%10, 10)

	return digitCheck
}
