package main

import (
	"fmt"
	"issue_2/output"
	"os"
)

func readCurrentDir() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)

	files,err := os.ReadDir(dir)
	if err != nil{
		panic(err)
	}
	if len(os.Args) > 1 && os.Args[1] == "-a"{
		output.OutputAll(files)
		return
	}
	output.Output(files)
}

func main(){
	readCurrentDir()
}
