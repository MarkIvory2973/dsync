package logs

import (
	"fmt"
	"os"
)

func Info(scope string, message string) {
	fmt.Fprintf(os.Stdout, "info(%s): %s\n", scope, message)
}

func Warning(scope string, message string, err error) {
	if err != nil {
		message = fmt.Sprintf("%s: %v", message, err)
	}

	fmt.Fprintf(os.Stderr, "warning(%s): %s\n", scope, message)
}

func Fatal(scope string, message string, err error) {
	if err != nil {
		message = fmt.Sprintf("%s: %v", message, err)
	}

	fmt.Fprintf(os.Stderr, "fatal(%s): %s\n", scope, message)
}
