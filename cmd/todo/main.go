package main

import (
	"bufio"
	"fmt"
	"flag"
	"github.com/samEscom/my_task/core"
	"os"
	"io"
	"strings"
	"errors"
)


const (
	dataFile = ".dataTodo.json"
)

func main(){
	add := flag.Bool("add", false, "can add new task to do")
	complete := flag.Int("complete", 0, "mark a task as completed")
	deleted := flag.Int("delete", 0, "delete a task")
	list := flag.Bool("list", false, "list of all tasks")

	flag.Parse()

	todos := &todo.Todos{}

	err := todos.Load(dataFile)


	if err != nil{
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch{
		case *add:
			task, err := getInput(os.Stdin, flag.Args()...)
			if err != nil{
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			todos.Add(task)
			err = todos.Store(dataFile)
			if err != nil{
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		case *complete > 0:
			err := todos.Complete(*complete)
			if err != nil{
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			
			err = todos.Store(dataFile)
			if err != nil{
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		case *deleted > 0:
			err := todos.Delete(*deleted)
			if err != nil{
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			
			err = todos.Store(dataFile)
			if err != nil{
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		case *list:
			todos.PrintTasks()
		default:
			fmt.Println(os.Stdout, "invalid a command")
			os.Exit(0) 		
	}
}


func getInput(r io.Reader, args... string) (string, error){

	if len(args) > 0{
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil{
		return "", err
	}

	if len(scanner.Text()) == 0{
		return "", errors.New("empty task, not allowed")
	}

	return scanner.Text(), nil

}