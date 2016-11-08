package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}

func memcp( in []byte, out[]byte, count int ) {
	i := 0
	for i < count {
		switch (count-i)/8 {
			default:
				fallthrough
			case 0:
				out[i] = in[i]
				i++
				fallthrough
			case 7:
				out[i] = in[i]
				i++
				fallthrough
			case 6:
				out[i] = in[i]
				i++
				fallthrough
			case 5:
				out[i] = in[i]
				i++
				fallthrough
			case 4:
				out[i] = in[i]
				i++
				fallthrough
			case 3:
				out[i] = in[i]
				i++
				fallthrough
			case 2:
				out[i] = in[i]
				i++
				fallthrough
			case 1:
				out[i] = in[i]
				i++
				break
		}
	}
}