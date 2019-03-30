package main

import (
	_ "net/http/pprof"
	"net/http"
	"log"
	"sync"
	"time"
	"io/ioutil"
	"fmt"
)

func main() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)

	for i := 0; i < 100; i++ {
		go doTask(&waitGroup)
	}

	go func() {

		log.Println(http.ListenAndServe("localhost:8082", nil))

	}()

	waitGroup.Wait()
	time.Sleep(3 * time.Second)
}

func doHttpRequest() (int) {

	result := []string{"request "}
	resp, err := http.Get("http://www.alibaba.com")
	if err != nil {

	}

	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {

	}
	result = append(result,string(bytes))
	return len(bytes)
}

func doTask(waitGroup *sync.WaitGroup) {

	for {

		for i:= 0 ;i<10000;i++{
			Add("hello world ! just do it!")
			len := doHttpRequest()

			Add(fmt.Sprintf("response length %v",len))
		}

		time.Sleep(3 * time.Second)
	}
	waitGroup.Done()
}

var datas []string

func Add(str string) string {
	data := []byte(str)
	dataStr := string(data)
	datas = append(datas, dataStr)
	return dataStr
}
