package app

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "3001"

func main() {
	fmt.Println("Started Application")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Welcome to k8s best manifests practices!")
	})

	fmt.Printf("Starting server at port %v\n", PORT)
	if err := http.ListenAndServe("localhost:"+PORT, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Exited Application")
}
