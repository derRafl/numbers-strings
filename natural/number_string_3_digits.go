package natural


// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {

	zahleneinsbisneuzehn := []string{"", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn"}

	zahlnull := "null"

	if n == 0 {
		return zahlnull
	}

	if n > 0 && n < 20 {
		return zahleneinsbisneuzehn[n]
	}

	var stellen []int = Digit(n)

	//fmt.Println(stellen)

	if len(stellen) == 2 {
		str_zwei := DigitString10(stellen[1])
		str_eins := DigitString1(stellen[0])
		return str_eins + str_zwei
	}

	if len(stellen) == 3 {
		zehner :=  stellen[1]*10 + stellen[0]

		if zehner < 20 {
			str_drei := DigitString100(stellen[2])
			return str_drei + "und" + zahleneinsbisneuzehn[zehner]
		}
		str_drei := DigitString100(stellen[2])
		str_zwei := DigitString1(stellen[0])
		str_eins := DigitString10(stellen[1])
		return str_drei + str_zwei + str_eins
	}

	return "Wenn das hier erschein gibt es einen fehler!"

}

func Digit(n int) []int {
	result := []int{}

	for n != 0 {
		r := n % 10
		n = n / 10

		result = append(result, r)
	}

	//slices.Reverse(result)
	return result
}
