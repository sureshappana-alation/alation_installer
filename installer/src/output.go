package main

import (
	"fmt"
	_ "reflect"
	_ "runtime"

	. "github.com/logrusorgru/aurora"
)

func logAndShowMsg(msg string) {
	LOGGER.Info(msg)
	output := fmt.Sprintf(" %s %s", Black("\u2714"), msg)
	fmt.Println(output)
}

func logAndShowSuccess(check string) {
	tmplMsg := "[ %s %s ] %s"
	output := fmt.Sprintf(tmplMsg, Green("\u2714"), Green("Success"), check)
	logOutput := fmt.Sprintf(tmplMsg, "\u2714", "Success", check)
	LOGGER.Info(logOutput)
	fmt.Println(output)
}

func logAndShowError(failure string, reason string) {
	output := fmt.Sprintf("[ %s %s ] %s. Failure Reason => %v", Red("\u2716"), Red("Failed"), failure, Red(reason))
	LOGGER.Error(fmt.Sprintf("%s %s", failure, reason))
	fmt.Println(output)
}
