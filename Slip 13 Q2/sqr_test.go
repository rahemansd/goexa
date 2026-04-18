package main

import "testing"

func Square(num int) int {
    return num * num
}

func BenchmarkSquare(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Square(10) 
    }
}


//to run -> go test -bench=.
