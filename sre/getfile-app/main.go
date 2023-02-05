package main

//This application gets the file from backend.
//Author: Nitin Jilla

import (
        "os"
        "fmt"
        "log"
        "time"
        "context"
        "net/http"
        "io/ioutil"
        "os/signal"
        "github.com/gorilla/mux"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)


//Displays a file chosen at random
func getRandomFile (w http.ResponseWriter, r *http.Request) {

        var fileExt string

        httpReqCounter.Inc()

        reqFile, err := http.Get("http://dummypdforpng-svc:3000")
        if err != nil {

                log.Println(err)
        }

        defer reqFile.Body.Close()

        respFile, err := ioutil.ReadAll(reqFile.Body)
        if err != nil {

                log.Println(err)
        }

        mType := http.DetectContentType(respFile)

        if mType == "application/pdf"{

                fileExt = "pdf"
        }else{

                fileExt = "png"
        }

        newFile := fmt.Sprintf("file.%s", fileExt)
        ioutil.WriteFile(newFile, respFile, 0644)

        http.ServeFile(w, r, newFile)
        log.Printf("A new file was served of type %s", mType)
}

//Provides application health
func healthCheck(w http.ResponseWriter, r *http.Request){

        fmt.Fprintf(w, "{status: OK}")                            //Not the correct way to do it
}

//Global variable which counts HTTP hits
var httpReqCounter = prometheus.NewCounter(
   prometheus.CounterOpts{
       Name: "http_request_count",
       Help: "No. of files downloaded.",
   } )



func main(){

        closeServer := make(chan os.Signal)                     //Creating a channel to catch the interrupt
        signal.Notify(closeServer, os.Interrupt)                //Notifies <-closeServer on Ctrl+C

        getFileServer := mux.NewRouter()                        //Create a new server

        httpSrv := &http.Server{
                Addr: ":3001",
                Handler: getFileServer,

        }

        getFileServer.HandleFunc("/", getRandomFile)            //Distributes a file chosen at random
        getFileServer.HandleFunc("/health", healthCheck)        //Endpoint for healthcheck
        getFileServer.Handle("/metrics", promhttp.Handler())

        prometheus.MustRegister(httpReqCounter)                 //Custom metrics for no. of hit received


        //Run our server in a goroutine so that it doesn't block
        go func(){
        err := httpSrv.ListenAndServe()                         //HTTP server runs on port 3001
        if err != nil {

                log.Println(err)
        }
        }()


        <-closeServer
        log.Println("Application will be terminated. Waiting for all open connections to close...")

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

        defer cancel()

        err := httpSrv.Shutdown(ctx);

        if err != nil{
                log.Println(err)
        }

        log.Println("Application was terminated.")

}

