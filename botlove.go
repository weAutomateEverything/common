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
	msg := fmt.Sprintf("ERROR: %v -  %v", os.Getenv("AWS_LAMBDA_FUNCTION_NAME"), err.Error())
	log.Printf(msg)
	resp, err := http.Post("http://api.carddevops.co.za/telegram-text/1083240170", "application/text", strings.NewReader(msg))
	if err == nil {
		resp.Body.Close()
	} else {
		log.Printf("HAL ERROR: %v", err.Error())
	}
	debug.PrintStack()
}

