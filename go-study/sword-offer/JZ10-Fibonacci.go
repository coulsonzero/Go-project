package main

/**
 * Q: JZ10 斐波那契数列
 *        { 1 				  x=1,2 }
 * fib(x)={                         }
 *		  { fib(x−1)+fib(x−2) x>2   }
 * input : 4
 * output: 3
 * explain: fib(1)=1,fib(2)=1,fib(3)=fib(3-1)+fib(3-2)=2,fib(4)=fib(4-1)+fib(4-2)=3
 */

func Fibonacci(n int) int {
	// write code here
	if n <= 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
