package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type IntVec99 [9][9]int

func (iv *IntVec99) Load(filepath string) {
    data, err := ioutil.ReadFile(filepath)
    if err != nil {
        fmt.Println(err)
        return
    }

    err = json.Unmarshal(data, iv)
    if err != nil {
        fmt.Println(err)
        return
    }
}

func (iv *IntVec99) Dump(filepath string) {
    data, err := json.Marshal(iv)
    if err != nil {
        fmt.Println(err)
        return
    }

    err = ioutil.WriteFile(filepath, data, 0664)
    if err != nil {
        fmt.Println(err)
        return
    }
}
