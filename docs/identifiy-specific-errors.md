[[toc]]

# 要解决的问题

* 如何识别具体的错误值

# 解决方案

* 类型不同&类型范围确定，可以通过判断错误值所属类型，尝试确定错误值范围
* 类型相同&值范围确定，可以通过错误值判等，尝试确定具体的错误值
* 类型不同&类型范围不确定，通过错误的字符串信息通过字符串判等、正则匹配等，尝试对错误值进行判断

# 解决方案案例 -- Go语言错误值判定

Go语言支持多返回值，对于错误处理有一套近乎通用的解决方案方式，即：
* 函数声明时，在结果列表的最后声明一个error类型的结果，以封装函数执行过程中的错误
* 函数调用时，先判断结果列表中最后一个结果的值是否为nil，以判断函数执行过程是否正常

然而，error类型本质上仅是Go语言自建的一个接口类型，其声明中也仅定义了一个无任何参数声明、结果声明有且仅有一个string类型的Error方法。如果想要进一步对错误进行有效处理，就必须实现错误值的准确识别！

## 类型不同&类型范围确定
* 类型断言表达式
```
err := writeFile(file, "new content")
if err != nil{
    if errRealVal, ok := err.(*FileReadError); ok{
        fmt.Println(fmt.Sprintf("FileReadError: %+v", errRealVal))
    }else if errRealVal, ok := err.(*FileWriteError); ok{
        fmt.Println(fmt.Sprintf("FileWriteError: %+v", errRealVal))
    }else{
        fmt.Println(fmt.Sprintf("FileError: %+v", errRealVal))
    }
}
```
* switch语句
```
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
```
源代码路径：../../identify-specific-error/same_type&known_type_range/demo.go

## 类型相同&值范围确定
* 判等表达式
```
err := validateFile(file)
if err != nil{
    if err == &FileReadError{
        fmt.Println(fmt.Sprintf("FileError: %+v", FileReadError))
    }else if err == &FileWriteError {
        fmt.Println(fmt.Sprintf("FileError: %+v", FileWriteError))
    }else{
        fmt.Println(fmt.Sprintf("FileError"))
    }
}
```
* switch语句
```
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
```
源代码路径：../../identify-specific-error/same_type&known_value_range/demo.go

## 类型不同&类型范围不确定
* 字符串匹配
```
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
```
源代码路径：../../identify-specific-error/diff_type&unknown_type_range/demo.go
