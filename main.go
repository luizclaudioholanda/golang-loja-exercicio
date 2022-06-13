package main

import (
	"net/http"

	"github.com/luizclaudioholanda/loja/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8888", nil)
}
