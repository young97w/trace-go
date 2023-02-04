package main

import "trace-go"

func foo() {
	defer trace.Trace()()
	bar()
}

func bar() {
	defer trace.Trace()()

}

func main() {
	defer trace.Trace()()
	foo()
}
