//UNIT TEST FOR DUMMY PDF OR PNG

//Author: Nitin Jilla


package main

import (
	"testing"
	"net/http"
	"io"
	"os"
	"fmt"
)

func TestMain(t *testing.T){
	

	//Send a GET request

	out, err := http.Get("http://192.168.8.124:4000")
	
	if err != nil{
	
		t.Error(err)
		t.FailNow()
	} 

	defer out.Body.Close()

	//Create the file

	dwFile, err := os.Create("./download/newfile")
	
	if err != nil{

                t.Error(err)
		t.FailNow()
        }
	
	defer dwFile.Close()
	
	//Copy content to that file

	_, err = io.Copy(dwFile, out.Body)
 
	if err != nil{
		
		t.Error(err)
		t.FailNow()
	}

	fmt.Println("File has been downloaded")
}
