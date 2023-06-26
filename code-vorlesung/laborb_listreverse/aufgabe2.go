// Thema Aufgabe: Listen
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe2

/* Aufgabenstellung:
 * In der Datei linkedlist.go  ist ein Datentyp für eine verkettete Liste definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode Reverse().
 * Die Methode soll die Liste umdrehen und den neuen Head liefern.
 * D.h. die Methode soll die Elemente erhalten, die Pointer aber so umbiegen, dass
 * die Liste umgedreht ist.
 *
 * Anmerkung: Bei dieser Aufgabe würde in der Bewertung auch geprüft werden,
 * ob wirklich nur Pointer verdreht wurden, d.h. ob die Adressen der Elemente noch
 * die gleichen sind. Diese Forderung ist aber relativ leicht einzuhalten,
 * indem man einfach keine neuen Elemente erzeugt ;-).
 */

// Dreht die Liste um und liefert den neuen Kopf.
func (list *LinkedList) Reverse() *LinkedList {

	// Rekursionsanker: Die leere Liste muss nicht umgedreht werden.
	if list.Empty() {
		return list
	}

	// 2. Rekursionsanker: Auch eine Liste mit nur einem Element muss nicht umgedreht werden.
	if list.Next.Empty() {
		return list
	}

	// Warum 2 Rekursionsanker?
	// Weil wir unten im Rekursionsschritt die Liste zerlegen und der hintere Teil
	// dann immer noch Länge >= 1 haben muss, damit alles funktioniert.

	// Wir merken uns die beiden ersten Elemente.
	A := list
	B := list.Next

	// Wir drehen alles ab B um.
	C := B.Reverse()

	// Wir hängen A zwischen B und das Ende.
	// Wir wissen schon, dass B jetzt das letzte Element des umgdrehten Rests ist.
	A.Next = B.Next
	B.Next = A

	return C
}
