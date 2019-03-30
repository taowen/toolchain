package main

import (
	"fmt"
	"strings"
)

type File struct{
	FilePath string
	CanBeRead bool
	CanBeWrite bool
}

type FileError struct {
	Val string
	ErrMsg string
}

func (f *FileError)Error()string{
	return fmt.Sprintf("File has error: %s",f.ErrMsg)
}

func validateFile(file File) error{
	if file.CanBeRead == false{
		return &FileError{file.FilePath, "File read failure!"}
	}else if file.CanBeWrite == false{
		return &FileError{file.FilePath, "File write failure!"}
	}

	return nil
}

func main(){
	// init
	file := File{"filepath", true, false}

	// 错误信息匹配
	err := validateFile(file)
	if err != nil{
		errMsg := err.Error()
		if strings.Contains(errMsg,"read"){
			fmt.Printf("file read failure: %s\n", errMsg)
		}else if strings.Contains(errMsg,"write"){
			fmt.Printf("file write failure: %s\n", errMsg)
		}else {
			fmt.Printf("file has failure: %s\n", errMsg)
		}
	}

	// 错误信息匹配
	file = File{"filepath", false, true}
	err = validateFile(file)
	if err != nil{
		errMsg := err.Error()
		if strings.Contains(errMsg,"read"){
			fmt.Printf("file read failure: %s\n", errMsg)
		}else if strings.Contains(errMsg,"write"){
			fmt.Printf("file write failure: %s\n", errMsg)
		}else {
			fmt.Printf("file has failure: %s\n", errMsg)
		}
	}
}