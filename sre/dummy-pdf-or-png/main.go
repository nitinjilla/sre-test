package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)



func serveRandomFile(w http.ResponseWriter, r *http.Request) {
	rnd := rand.Intn(11)

	fmt.Println(rnd)

	if rnd < 5 {
		http.ServeFile(w, r, "./dummy.png")
		return
	}

	if rnd >=5 && rnd < 9 {
		http.ServeFile(w, r, "./dummy.pdf")
		return
	}

	http.ServeFile(w, r, "./corrupt-dummy.pdf")


}

func main() {
	
	rand.Seed(time.Now().UnixNano())
	
	http.HandleFunc("/", serveRandomFile)
	
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}

}

