/*Le programme suivant doit être modifié de tel qu’il utilise un « slice » de
taille arbitraire au lieu d’un tableau fixe.
La taille soumit par l’utilisateur (variable sz) doit remplacer la valeur 32
codé en dur de la taille du tableau. Modifiez la fonction main, ainsi que
les autres fonctions si nécessaire.
Soumettez votre solution en un fichier go. */

package main

import "fmt"
import "runtime"
import "math"
import "math/rand"

type Series struct {
	a, b float64
}

func fourier(c []Series, t, TP int, out chan float64) {
	res := c[0].a
	for n := 1; n < len(c); n++ {
		res += c[n].a*math.Sin(2.0*math.Pi*float64(t)/float64(TP)) + c[n].b*math.Cos(2.0*math.Pi*float64(t)/float64(TP))
	}
	out <- res
	return
}

func receive(data chan float64, TP int) (endOfCalc chan bool) {
	endOfCalc = make(chan bool)
	go func() {
	for t := 0; t < TP; t++ {
		f := <-data
		fmt.Printf("%f ", f)
	}
	fmt.Println()
	endOfCalc <- true
	}()
	return
}

func main() {
	runtime.GOMAXPROCS(3)

	fmt.Print("Enter length: ")
	var N int
	fmt.Scanf("%d", &N)
	data := make(chan float64)
	defer close(data)
	var c []Series
	c = make([]Series, N)
	TP := 4

	for t := 0; t < TP; t++ {
		for k := 0; k < N; k++ {
			c[k].a = rand.Float64()
			c[k].b = rand.Float64()
		}
		go fourier(c[:], t, TP, data)
	}

	endOfCalc := receive(data, TP)
	defer close(endOfCalc)
	eoc := <-endOfCalc
	if eoc {
		fmt.Println("All done")
	}
}
