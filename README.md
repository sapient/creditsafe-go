# Creditsafe Company Information API

## Overview

The Creditsafe Company Information API is a Go-based API that allows you to retrieve company credit ratings and other relevant company information from the Creditsafe website. This API simplifies the process of accessing credit information, enabling you to integrate it into your applications or services.

## Features

- Retrieve company credit ratings.
- Access detailed company information, including financial data.
- Built with Go for reliability and efficiency.

## Prerequisites

Before you can use this API, ensure you have the following prerequisites:

- Go installed on your system.
- Username and password from Creditsafe (sign up at https://www.creditsafe.com).

## Installation

```sh
go mod add github.com/sapient/creditsafe
```

## Example Usage

```go
package main

import (
	"fmt"
	"github.com/sapient/creditsafe"
)

func main() {
	csClient := creditsafe.New(
		creditsafe.SetUsername("username"),
		creditsafe.SetPassword("password"),
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
```