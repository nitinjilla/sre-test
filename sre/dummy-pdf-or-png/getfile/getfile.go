package main

import (
        "net/http"
        "io"
        "os"
        "log"
        "fmt"
        "github.com/gabriel-vasile/mimetype"
)

func rootFunc(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Visit http://testerapp.com:3001/yourfile to download your file to download your file")
}

func getRandomFile(w http.ResponseWriter, r *http.Request) {

        //Logs
        logfile, _ := os.OpenFile("/var/log/dummy-pdf-or-png.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        log.SetOutput(logfile)
        infoLogger := log.New(logfile, "INFO: ", log.LstdFlags|log.Lshortfile)

        //Send a GET request
        out, err := http.Get("http://192.168.8.119:3000")
        if err != nil{
                infoLogger.Println(err)
        }

        defer out.Body.Close()

        //Create the file
        dwFile, err := os.Create("/download/dummy")
        if err != nil{
                infoLogger.Println(err)
        }

        defer dwFile.Close()

        //Copy content to that file
        _, err = io.Copy(dwFile, out.Body)
        if err != nil{
                infoLogger.Println(err)
        }

        mType, err := mimetype.DetectFile("/download/dummy")
        infoLogger.Printf("A document of type %v was downloaded. MIME: %v.", mType.Extension(), mType.String())
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "{health: OK}")
}


func main(){
        http.HandleFunc("/", rootFunc)
        http.HandleFunc("/yourfile", getRandomFile)
        http.HandleFunc("/healthcheck", healthCheck)
        err := http.ListenAndServe(":3001", nil)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Your file has been downloaded.")
}

