package main

import (
	"net/http"
	"fmt"
)


func hello(w http.ResponseWriter, req *http.Request) {
	res := GetUrlResp(req.URL.Path)
	genResponse(res, w)
}


func main(){
	go ConfLoaderTimer("./conf.json")
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Printf("ListenAndServe on 8001 err:%s",err.Error())
	}
}

