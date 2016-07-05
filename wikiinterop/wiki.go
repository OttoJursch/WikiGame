package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
)

func main(){
    resp, err := http.Get("https://en.wikipedia.org/w/api.php?action=query&format=json&prop=links%7Clinkshere&titles=Albert+Einstein&pllimit=5000")
    if err != nil{
        fmt.Println("ERROR")
    }
    body, _ := ioutil.ReadAll(resp.Body)
    json := string(body)
    fmt.Println(json)
    resp.Body.Close()
}
