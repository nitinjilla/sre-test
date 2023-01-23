//App: Dummy PDF OR PNG

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"os"
)


func main() {
	
	//HTTP function
	rand.Seed(time.Now().UnixNano())
	
	http.HandleFunc("/", serveRandomFile)
	
	err := http.ListenAndServe(":4000", nil)

	//Logging for application
	file, _ := os.OpenFile("/var/log/file-sharer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(file)
	infoLogger := log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile)	

	infoLogger.Println("A file was downloaded.")	

	if err != nil {
		log.Fatal(err)
	}

	file.Close()

}


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
