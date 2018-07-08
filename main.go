package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	goPrime := func(v []js.Value) {
		js.Global().Set("output", js.ValueOf(isPrime(int64(v[0].Int()))))
		js.Global().Call("done")
	}
	js.Global().Set("goPrime", js.NewCallback(goPrime))
	<-c
}

/*
 function is_prime(n)
     if n ≤ 1
        return false
     else if n ≤ 3
        return true
     else if n mod 2 = 0 or n mod 3 = 0
        return false
     let i ← 5
     while i * i ≤ n
        if n mod i = 0 or n mod (i + 2) = 0
            return false
        i ← i + 6
		 return true
*/

func isPrime(n int64) bool {
	fmt.Println(n)
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if (n%2 == 0) || (n%3 == 0) {
		return false
	}
	var i int64 = 5
	for i*i <= n {
		if (n%i == 0) || n%(i+2) == 0 {
			return false
		}
		i = i + 6
	}
	return true
}
