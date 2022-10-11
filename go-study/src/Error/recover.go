package main

// recover()：捕获所有可能发生的异常，并将内部异常转换为错误处理, 必须在defer()函数中直接调用recover()
// recover()函数捕获的是父一级调用函数栈帧的异常
// panic(): 将异常抛出为相应的错误信息

func main() {
	defer func() {
		if r := recover(); r != nil {
			// r.(type): string, runtime.Error, error, ...
			// err = ...
			panic(r)
		}
	}()

}
