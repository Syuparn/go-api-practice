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
	personRepository := api.NewPersonRepository()

	switch os.Args[1] {
	case "create":
		age, name := argsForCreate()
		con := controller.NewCreateController(personFactory)
		err := con.Create(age, name)
		if err != nil {
			fmt.Println(err.Error())
		}
	case "read":
		con := controller.NewReadController(personRepository)
		err := con.Read()
		if err != nil {
			fmt.Println(err.Error())
		}
	case "update":
		age, id, name := argsForUpdate()
		con := controller.NewUpdateController(personRepository)
		err := con.Update(age, id, name)
		if err != nil {
			fmt.Println(err.Error())
		}
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

func argsForUpdate() (int, string, string) {
	f := flag.NewFlagSet("update", flag.ExitOnError)
	age := f.Int("age", 0, "age of person")
	name := f.String("name", "Anon", "name of person")
	id := f.String("id", "00000000-0000-0000-0000-000000000000", "uuid of person")
	f.Parse(os.Args[2:])

	return *age, *id, *name
}
