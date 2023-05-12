package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func ResolveStatusCode(input string) (string, error) {
	code, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("%s is not a valid input", input)
	}

	txt := http.StatusText(code)
	if txt == "" {
		return "", fmt.Errorf("%d is not a valid status code", code)
	}

	return fmt.Sprintf("%d: %s", code, txt), nil
}