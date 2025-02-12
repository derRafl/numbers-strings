package natural

// NumberStringBelow20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringBelow20(n int) string {

	zahleneinsbisneuzehn := []string{"", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn"}

	zahlnull := "null"

	if n == 0 {
		return zahlnull
	}
	return zahleneinsbisneuzehn[n]
}
