package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getName() (string){
	
	dat, err := ioutil.ReadFile("writer.dat")
    check(err)
    return string(dat)
}


