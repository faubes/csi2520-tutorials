/* On dispose d'un tableau comme par exemple le suivant:

var tableau = [7]int{3, 4, 8, 9, 5, 2, 7}

et on voudrait pouvoir trouver les éléments minimum et maximum
dans ce tableau et les déplacer au début du tableau.
Après cette opération, le contenu du tableau sera donc comme suit:

{2, 9, 8, 4, 5, 3, 7}

L'ordre des autres éléments n'a pas d'importance pourvu que le premier élément soit le minimum et le second soit le maximum.

Écrire une fonction Go permettant d'effectuer cette opération.
Évidemment, cette fonction doit pouvoir fonctionner pour n'importe quel
tableau (ou slice).
*/

package main

import "fmt"

func moveMinMax(a []int) []int {
	b := make([]int, len(a))
  copy(b, a)
	if len(a) < 2 {
		return b
	}
	var (
		min int = b[0]
		max int = b[0]
	)
	for i, v := range b {
		if i == 0 {
			continue
		}
		if v < min {
			min = v
			b[0], b[i] = min, b[0] // swap
		}
		if v > max {
			max = v
			b[1], b[i] = max, b[1] // swap
		}
	}
	return b
}

func moveAndPrint(a []int) {
  fmt.Println("Before moveMinMax", a)
  b := moveMinMax(a)
  fmt.Println("After moveMinMax", b)
}

func main() {
	var t1 = [7]int{3, 4, 8, 9, 5, 2, 7}
  var t2 = [7]int{3, 1, 1, 1, 1, 1, 1}
  var t3 = [7]int{1, 9, 2, 3, 4, 5, 6}
  var t4 = [7]int{9, 1, 2, 3, 4, 5, 6}
  moveAndPrint(t1[:])
  moveAndPrint(t2[:])
  moveAndPrint(t3[:])
  moveAndPrint(t4[:])
}
