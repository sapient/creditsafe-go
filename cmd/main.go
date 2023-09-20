package main

import (
	"flag"
	"fmt"

	"github.com/sapient/creditsafe"
)

func main() {
	username := flag.String("username", "", "Creditsafe username")
	password := flag.String("password", "", "Creditsafe password")
	companyId := flag.String("company-id", "", "Creditsafe company id")
	flag.Parse()

	csClient := creditsafe.New(
		creditsafe.SetUsername(*username),
		creditsafe.SetPassword(*password),
		creditsafe.SetSandbox(true),
	)

	err := csClient.Login()
	if err != nil {
		fmt.Println(err)
	}

	report, err := csClient.GetCompanyReport(*companyId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", report)
}
