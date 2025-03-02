package aufgabe2

/* AUFGABENSTELLUNG: Vervollständigen Sie die Funktion FilterVowels.
 * ERREICHBARE PUNKTE: 10
 */

// FilterVowels erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Vokale entfernt werden.
// Anmerkung: Es müssen nur Kleinbuchstaben beachtet werden.
func FilterVowels(s string) string {
	result := ""
	// TODO

	for _, el := range s {
		if el != 'a' && el != 'e' && el != 'i' && el != 'o' && el != 'u' {

			result += string(el)
		}
	}
	return result
}
