package natural

//import "fmt"

// NumberString6Digits erwartet eine Zahl 0 <= n <= 999999 und liefert den zugehörigen String.
func NumberString6Digits(n int) string {

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
		str_zwei := DigitString10(stellen[0])
		str_eins := DigitString1(stellen[1])
		return str_eins + str_zwei
	}

	if len(stellen) == 3 {
		zehner :=  stellen[1]*10 + stellen[0]

		if zehner < 20 {
			str_drei := DigitString100(stellen[2])
			return str_drei + "und" + zahleneinsbisneuzehn[zehner]
		}
		str_drei := DigitString100(stellen[0])
		str_zwei := DigitString1(stellen[2])
		str_eins := DigitString10(stellen[1])
		return str_drei + str_zwei + str_eins
	}

	if len(stellen) > 3 {
		tausender :=  stellen[4]*10 + stellen[3]
		var tausender_schriftlich string  = ""
		//fmt.Println(tausender)

		if tausender < 20{
			tausender_schriftlich = zahleneinsbisneuzehn[tausender]
		}else{
			str_vier := DigitString1(stellen[3])
			str_fünf := DigitString10(stellen[4])
			tausender_schriftlich = str_vier + str_fünf

		}

		
		if len(stellen) > 5 {
			str_sechs := DigitString100(stellen[5])
			tausender_schriftlich = str_sechs + tausender_schriftlich
		}
			
		tausender_schriftlich = tausender_schriftlich+"tausend"

		str_drei := DigitString100(stellen[2])

		zehner :=  stellen[1]*10 + stellen[0]

		if zehner < 20 {
			str_drei = str_drei + zahleneinsbisneuzehn[zehner]
		}
		
		str_zwei := DigitString1(stellen[0])
		str_eins := DigitString10(stellen[1])
		return tausender_schriftlich + str_drei + str_zwei + str_eins
	}

	return "Wenn das hier erschein gibt es einen fehler!"

}
