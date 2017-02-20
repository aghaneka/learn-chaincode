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
	pwd, _ := os.Getwd()
	fmt.Printf(pwf)
	dat, err := ioutil.ReadFile(pwd+"/writer.dat")
    check(err)
    return string(dat)
}


