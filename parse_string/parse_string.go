package parse_string

// ParseString erwartet einen String, der entweder eine Dezimal- oder eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseString(s string) int {
	// TODO

	hex := false
	for _, n := range s {
		if ParseDigit(string(n)) > 9 {
			hex = true
		}

		if hex {
			return ParseStringHexadecimal(s)
		}
	}
	return ParseStringDecimal(s)
}
