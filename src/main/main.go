package main

import (
	"fmt"
	"math"
	"net/http"
)

func efetuaConta() float64 {
	soma := 0.00
	for i := 0.00; i < 400000000; i++ {
		soma = soma + math.Sqrt(i)
	}
	return soma
}

func exibe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A soma é %f", efetuaConta())
}

func main() {
	http.HandleFunc("/", exibe)
	http.ListenAndServe(":8000", nil)
}
