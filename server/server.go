package main

import (
	"fmt"
	"net/http"
)

type serverhandler struct {
	mux *http.ServeMux
}

func (sh *serverhandler) run(p string) {
	var port string = p

	server := http.Server{
		Addr:    port,
		Handler: sh.mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("server err")
	}

}

func newServerInit() *serverhandler {
	fmt.Println("server ready")

	mux := http.NewServeMux()

	return &serverhandler{
		mux: mux,
	}

}
