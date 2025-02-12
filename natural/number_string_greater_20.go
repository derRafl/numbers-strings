package natural

// NumberStringGreater20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringGreater20(n int) string {

	var stellen []int = Digit(n)

	if len(stellen) == 2 {
		str_zwei := DigitString10(stellen[1])
		str_eins := DigitString1(stellen[0])
		return str_eins + str_zwei
	}

	if len(stellen) == 3 {
		zehner := stellen[1]*10 + stellen[0]

		if zehner == 0 {
			return DigitString100(stellen[2])
		}

		if zehner < 20 {
			str_drei := DigitString100(stellen[2])
			return str_drei + NumberStringBelow20(zehner)
		}

		str_drei := DigitString100(stellen[2])
		str_zwei := DigitString1(stellen[0])
		str_eins := DigitString10(stellen[1])
		return str_drei + str_zwei + str_eins
	}

	return "Wenn du das liest, gibt es einen fehler!"
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
