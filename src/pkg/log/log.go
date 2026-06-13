package log

import (
	"fmt"
)

func Note(message string) {
	fmt.Printf("\033[1mnote: %s\033[0m\n", message)
}

func Warning(message string, err error) {
	if err != nil {
		fmt.Printf("\033[1;33mwarning:\033[1;37m %s: %v\033[0m\n", message, err)
	} else {
		fmt.Printf("\033[1;33mwarning:\033[1;37m %s\033[0m\n", message)
	}

}

func Fatal(message string, err error) {
	if err != nil {
		fmt.Printf("\033[1;31mfatal:\033[1;37m %s: %v\033[0m\n", message, err)
	} else {
		fmt.Printf("\033[1;31mfatal:\033[1;37m %s\033[0m\n", message)
	}
}
