package main

import (
	"fmt"
	"flag"
	"github.com/samEscom/my_task/core"
	"os"
)


const (
	dataFile = ".dataTodo.json"
)

func main(){
	add := flag.Bool("add", false, "can add new task to do")

	flag.Parse()

	todos := &todo.Todos{}

	err := todos.Load(dataFile)


	if err != nil{
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch{
		case *add:
			todos.Add("sample")
			err := todos.Store(dataFile)
			if err != nil{
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			} 
		default:
			fmt.Println(os.Stdout, "invalida command")
			os.Exit(0) 		
	}
}
