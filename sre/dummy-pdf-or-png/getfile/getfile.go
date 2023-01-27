//getfile.go

//Author: Nitin Jilla

package main

import (
	"net/http"
	"io"
	"os"
	"os/signal"
//	"fmt"
	"log"
	"time"
	"context"
)

func main(){

	
	//Catch the interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	//Microservice logs
	logfile, _ := os.OpenFile("/var/log/dummy-pdf-or-png.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(logfile)
	infoLogger := log.New(logfile, "INFO: ", log.LstdFlags|log.Lshortfile)
	

	//Graceful shutdown of Application
	srv := http.Server{}
	srv.ListenAndServe()
	
	go func(){
		<-c
		time.Sleep(5 * time.Second)
		infoLogger.Println("Application will be termintated!")
		if err := srv.Shutdown(context.Background()); err != nil{

		infoLogger.Println("Application terminated")

		}
	}()

	//Send a GET request
	out, err := http.Get("http://192.168.8.116:3000")
	
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

	infoLogger.Println("File has been downloaded")
}
