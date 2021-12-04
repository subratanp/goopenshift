package main

import (
	"apifuncs"
	"fmt"
	"net/http"
)

type Test struct {
}

func main() {
	sm := http.NewServeMux()
	apsm := apifuncs.Apifunc{}
	sm.Handle("/getData", &apsm)
	server := http.Server{
		Addr:    ":9090",
		Handler: sm,
	}

	server.ListenAndServe()
	fmt.Println("This is a great thing")

}
