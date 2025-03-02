package parse_string

// ParseDigit erwartet einen String von 0 bis 9 oder A bis F und liefert die zugehörige Zahl.
// Dabei gilt A=10, B=11, ..., F=15.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseDigit(digit string) int {
	// TODO
	switch digit {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "A":
		return 10
	case "B":
		return 11
	case "C":
		return 12
	case "D":
		return 13
	case "E":
		return 14
	case "F":
		return 15
	default:
		return -1
	}
}
