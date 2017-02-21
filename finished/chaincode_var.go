package main

import (
    "net/http"
//    "fmt"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//func main(){
//	fmt.Printf("Name : "+getName())
//}

func getName() (string){
	resp, err := http.Get("http://52.220.79.230:8080/write.html")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
    return string(body)
    
}


