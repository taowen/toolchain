package main

import (
	"fmt"
)

type File struct{
	FilePath string
	CanBeRead bool
	CanBeWrite bool
}

type FileReadError struct {
	Val string
	CanBeRead bool
}

func (f *FileReadError)Error()string{
	return fmt.Sprintf("File can not be read: %s",f.Val)
}

type FileWriteError struct {
	Val string
	CanBeWrite bool
}

func (f *FileWriteError)Error()string{
	return fmt.Sprintf("File can not be write: %s",f.Val)
}

func readFile(file File)(string, error){
	if file.CanBeRead == false{
		return "", &FileReadError{file.FilePath, false}
	}

	// 读取文件...

	return "The content of your file", nil
}

func writeFile(file File, content string)error{
	if file.CanBeWrite == false{
		return &FileWriteError{file.FilePath, false}
	}

	// 写入文件...

	return  nil
}

func main(){
	// init
	file := File{"filepath", false, false}

	// 1.switch语句
	_, err := readFile(file)
	if err != nil{
		switch errRealType := err.(type) {
		case *FileReadError:
			fmt.Println(fmt.Sprintf("FileReadError: %+v", errRealType))
		case *FileWriteError:
			fmt.Println(fmt.Sprintf("FileWriteError: %+v", errRealType))
		default:
			fmt.Println(fmt.Sprintf("FileError: %+v", errRealType))
		}
	}

	// 2.类型断言表达式
	err = writeFile(file, "new content")
	if err != nil{
		if errRealVal, ok := err.(*FileReadError); ok{
			fmt.Println(fmt.Sprintf("FileReadError: %+v", errRealVal))
		}else if errRealVal, ok := err.(*FileWriteError); ok{
			fmt.Println(fmt.Sprintf("FileWriteError: %+v", errRealVal))
		}else{
			fmt.Println(fmt.Sprintf("FileError: %+v", errRealVal))
		}
	}
}