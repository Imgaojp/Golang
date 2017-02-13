package main

import "fmt"

const(
	i=1<<iota
	j=3<<iota
	k=4<<iota
	l
)

func main()  {
	fmt.Println(i,j,k,l)
}