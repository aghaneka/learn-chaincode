package main

import (
    "io/ioutil"
    "fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getName() (string){

	dat, err := ioutil.ReadFile("/root/writer.dat")
    check(err)
    return string(dat)
    
}


