package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE:clac command [arguments]")
	fmt.Println("-------------------------------")
}

func main() {
	args := os.Args
	//fmt.Println(args)
	//println(args[0])
	//println(args[1])
	//println(args[2])
	//println(args[3])
	//println(len(args))
	//fmt.Println(simplemath.Age)
	if args == nil || len(args) < 4 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add integer1 integer2")
			return
		}

		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])

		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add integer1 integer2")
			return
		}

		result := simplemath.Add(v1, v2)
		fmt.Println("Result: ", result)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
