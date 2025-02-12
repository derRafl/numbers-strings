package natural

// NumberString6Digits erwartet eine Zahl 0 <= n <= 999999 und liefert den zugehörigen String.
func NumberString6Digits(n int) string {

	//Szahleneinsbisneuzehn := []string{"", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn"}

	if n < 1000 {
		return NumberString3Digits(n)
	}

	stellen := Digit(n)

	uk := stellen[0] + stellen[1]*10 + stellen[2]*100
	ok := 0

	if len(stellen) == 4 {
		ok = stellen[3]
	}

	if len(stellen) == 5 {
		ok = stellen[3] + stellen[4]*10
	}

	if len(stellen) == 6 {
		ok = stellen[3] + stellen[4]*10 + stellen[5]*100
	}

	if uk == 0 {
		return NumberString3Digits(ok) + "tausend"

	}
	return NumberString3Digits(ok) + "tausend" + NumberString3Digits(uk)

	//return "Wenn das hier erschein gibt es einen fehler!"

}
