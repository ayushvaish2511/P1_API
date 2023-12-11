package main

import (
	"P1_API/helper"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("Start Server")

	routes := httprouter.New()

	routes.GET("/",func (w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
		fmt.Fprint(w, "Welcome Home")
	})

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}