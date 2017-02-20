package main

import (
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getName() (string){
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd+"/writer.dat")
    check(err)
    return string(dat)
}


