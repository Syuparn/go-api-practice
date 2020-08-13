package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Syuparn/go-api-practice/client/api"
	"github.com/Syuparn/go-api-practice/client/controller"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("sub command must be required. (create|read|update|delete)")
		return
	}

	personFactory := api.NewPersonFactory()

	switch os.Args[1] {
	case "create":
		age, name := argsForCreate()
		con := controller.NewCreateController(personFactory)
		con.Create(age, name)
	default:
		fmt.Println("sub command must be (create|read|update|delete)")
	}
}

func argsForCreate() (int, string) {
	f := flag.NewFlagSet("create", flag.ExitOnError)
	age := f.Int("age", 0, "age of person")
	name := f.String("name", "Anon", "name of person")
	f.Parse(os.Args[2:])

	return *age, *name
}
