package main

import "fmt"

func main() {
	//fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
	//fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}))
	//fmt.Println(exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}))
	//fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:1", "0:start:2", "0:end:3", "0:end:4", "0:end:5"}))
	fmt.Println(exclusiveTime(3, []string{"0:start:0", "1:start:2", "1:end:5", "2:start:6", "2:end:9", "0:end:12"}))
}
