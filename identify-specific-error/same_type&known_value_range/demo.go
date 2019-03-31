package main

import (
	"fmt"
)

var (
	FileReadError = FileError{
		FilePath: "the path",
		CanNotBeRead: true,
	}
	FileWriteError = FileError{
		FilePath: "the path",
		CanNotBeWrite: true,
	}
)

type File struct{
	FilePath string
	CanBeRead bool
	CanBeWrite bool
}

type FileError struct {
	FilePath string
	CanNotBeRead bool
	CanNotBeWrite bool
}

func (f *FileError)Error() string{
	return fmt.Sprintf("%+v", f)
}


func validateFile(file File) error{
	if file.CanBeRead == false{
		return &FileReadError
	}else if file.CanBeWrite == false{
		return &FileWriteError
	}

	return nil
}


func main(){
	// init
	file := File{"filepath", true, false}

	// 1.switch语句
	err := validateFile(file)
	if err != nil{
		switch err {
		case &FileReadError:
			fmt.Println(fmt.Sprintf("FileError: %+v", FileReadError))
		case &FileWriteError:
			fmt.Println(fmt.Sprintf("FileError: %+v", FileWriteError))
		default:
			fmt.Println(fmt.Sprintf("FileError"))
		}
	}

	// 2.判等表达式
	file = File{"filepath", false, true}
	err = validateFile(file)
	if err != nil{
		if err == &FileReadError{
			fmt.Println(fmt.Sprintf("FileError: %+v", FileReadError))
		}else if err == &FileWriteError {
			fmt.Println(fmt.Sprintf("FileError: %+v", FileWriteError))
		}else{
			fmt.Println(fmt.Sprintf("FileError"))
		}
	}
}