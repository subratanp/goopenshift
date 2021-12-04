package apifuncs

import (
	"fmt"
	"net/http"
)

type Apifunc struct {
}

func (pn *Apifunc) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
	fmt.Println("This is running well")
	wr.Write([]byte("This is working great"))

}
