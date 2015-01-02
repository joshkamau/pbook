package main

import (
	"fmt"
	"os"
	"path/filepath"
	"bitbucket.org/joshnet/pbook/model"
	"bitbucket.org/joshnet/pbook/dao"
	"log"
)

var help = fmt.Sprintf("usage : %s [all|find <name>|add <name> <phone number>]", filepath.Base(os.Args[0]))

func main(){
	if len(os.Args) == 1 {
		fmt.Println(help)
		os.Exit(0)
	}

	if len(os.Args) > 1 {
		command := os.Args[1]

		switch command {
		case "all":
			handleAll()
		case "find":
			handleFind(os.Args)
		case "add":
			handleAdd(os.Args)
		default:
			fmt.Println(help)
		}
	}
}

func displayContacts(contacts []*model.Contact){
	if len(contacts) > 0 {
		for _, c := range contacts {
			fmt.Printf("%d	%s	%s\n", c.Id, c.Name, c.PhoneNumber)
		}
	} else {
		fmt.Println("No Contacts")
	}
}

func handleAll() {
	contacts, err := dao.GetAll()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	displayContacts(contacts)
	os.Exit(0)
}

func handleFind(args []string) {
	if len(args) < 3 {
		contacts, err := dao.GetAll()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		displayContacts(contacts)
	}

	name := args[2]
	fmt.Printf("Showing results for: %s \n", name)
	contacts, err := dao.GetByName(name)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	displayContacts(contacts)
	os.Exit(0)
}

func handleAdd(args []string) {
	if len(args) < 4 {
		fmt.Println(help)
		os.Exit(0)
	}

	name := args[2]
	phoneBook := args[3]

	contact := model.Contact{Name:name, PhoneNumber:phoneBook}
	id, err := dao.SaveContact(&contact)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("Contact created. ID: %d\n", id)
 }
