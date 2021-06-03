package main

import (
	"log"
	"net/http"
	"os"
	"github.com/catmandude/nic-jackson-go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.Goodbye(l)

	sm := http.NewServeMux() //points to the endpoint /thing
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm) //specify our servemux
}
