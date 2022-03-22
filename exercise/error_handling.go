package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ""
	res, err := requiredValidation(input)
	if err != nil {
		fmt.Printf("%s \n", err.Error())
	} else {
		fmt.Println(res)
	}

}

func requiredValidation(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", RequiredError()
	}
	return input, nil
}
