package main

import (
    "io/ioutil"
    "os"
    "fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getName() (string){
	
	return "Geoff Boycott"
	/*
	pwd, _ := os.Getwd()
	fmt.Printf("dir : "+pwd)
	dat, err := ioutil.ReadFile(pwd+"/writer.dat")
    check(err)
    return string(dat)
    */
}


