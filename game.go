package main

import (
    "fmt"
    "net/http"
    "html/template"
)

func main(){
    http.HandleFunc("/", baseHandler)
    http.ListenAndServe(":8080", nil)
}

func baseHandler(writer http.ResponseWriter, r *http.Request){
    fmt.Println("testing: we are on port 8080")
    text := "HELLO WORLD"
    temp, err := template.New("test").Parse(text)
    fmt.Println(err)
    temp.Execute(writer, nil)
}
