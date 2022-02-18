package app

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "3001"

func main() {
	fmt.Println("Started Application")

	http.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service is live!")
	})

	fmt.Printf("Starting server at port %v\n", PORT)
	if err := http.ListenAndServe("localhost:"+PORT, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Exited Application")
}
