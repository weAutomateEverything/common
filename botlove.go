package common

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

func LogLambdaError(err error) {
	log.Printf("ERROR: %v -  %v", os.Getenv("AWS_LAMBDA_FUNCTION_NAME"), err.Error())
	resp, err := http.Post("http://api.carddevops.co.za/telegram/error", "application/text", strings.NewReader(fmt.Sprintf("*%v*\n\n%v", err.Error())))
	if err == nil {
		resp.Body.Close()
	} else {
		log.Printf("HAL ERROR: %v", err.Error())
	}
	debug.PrintStack()
}

