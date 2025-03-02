package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	// TODO
	seenfirst := false
	seenlast := false
	result := []string{}

	for _, el := range list {

		if el == first {
			seenfirst = true
		}

		if el == last {
			seenlast = true
		}

		if !seenfirst && seenlast {
			return []string{}
		}

		if seenfirst == seenlast && el != last {
			result = append(result, el)
		}

	}

	if !seenfirst && !seenlast {
		return []string{}
	}

	return result
}
