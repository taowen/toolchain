package main

import (
	"plugin"
	"fmt"
	"errors"
	"os"
)

const fnKey = "GetFunc"

func main() {

	p, err := plugin.Open("./pluginhello_v1.so")
	if err != nil {
		fmt.Println("error open plugin: ", err)
		os.Exit(-1)
	}
	fn, e := GetPluginKeyByName(p, fnKey)

	if e != nil {
		fmt.Printf("lookup error =%+v\n", e)
		os.Exit(-1)
	}

	result := fn("world")
	fmt.Println(result)

}

func GetPluginKeyByName(p *plugin.Plugin, symbolName string) (func(args ...interface{}) string, error) {
	if p == nil {

		return nil, errors.New("p为空指针.")
	}
	sym, err := p.Lookup(symbolName)
	if err != nil {
		return nil, err
	}
	fn, ok := sym.(func() func(args ...interface{}) string)
	if !ok {
		return nil, fmt.Errorf("symbol:%T ", sym)
	}
	if fn() == nil {
		return nil, errors.New("fn为空.")
	}
	return fn(), nil

}
