package trace

import "testing"

func a() {
	defer Trace()()
	b()
}

func b() {
	defer Trace()()
	c()
}

func c() {
	defer Trace()()
	d()
}

func d() {
	defer Trace()()
}

func TestTrace(t *testing.T) {
	a()
	//Outputs:
	//g[00004]:   ->trace-go.a
	//g[00004]:      ->trace-go.b
	//g[00004]:         ->trace-go.c
	//g[00004]:            ->trace-go.d
	//g[00004]:            <-trace-go.d
	//g[00004]:         <-trace-go.c
	//g[00004]:      <-trace-go.b
	//g[00004]:   <-trace-go.a
}
