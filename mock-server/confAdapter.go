package main

import (
	"sync"
	"io/ioutil"
	"encoding/json"
	"time"
	"fmt"
)

var UrlPatternMap = struct{
	sync.RWMutex
	m ResConf
}{m: make(map[string]*Response)}

func ConfLoader (filePath string) error {
	UrlPatternMap.Lock()
	defer UrlPatternMap.Unlock()
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &UrlPatternMap.m)
	if err != nil {
		return err
	}
	return nil
}

func ConfLoaderTimer(filePath string) {
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <- tick:
			err := ConfLoader(filePath)
			if err != nil {
				fmt.Printf("load conf err:%s\n",err.Error())
			}else {
				str, _ := json.Marshal(UrlPatternMap.m)
				fmt.Printf("load conf success:%s\n", str)
			}
		}
	}
}

func GetUrlResp(url string) *Response {
	UrlPatternMap.RLock()
	defer 	UrlPatternMap.RUnlock()
	if res,ok := UrlPatternMap.m[url]; ok {
		return res
	}

	return NewNotFoundUrlRes()
}
