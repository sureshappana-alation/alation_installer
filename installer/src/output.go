package main

import (
	"fmt"
	_ "reflect"
	_ "runtime"

	. "github.com/logrusorgru/aurora"
)

func print(messages ...interface{}) {
	fmt.Println(messages...)
}

func show_result(check string) {
	tmpl_msg := "[ %s %s ] %s"
	output := fmt.Sprintf(tmpl_msg, Green("\u2714"), Green("Success"), check)
	log_output := fmt.Sprintf(tmpl_msg, "\u2714", "Success", check)
	LOGGER.Info(log_output)
	print(output)
}

func show_error(check string, msg string) {
	tmpl_msg := "[ %s %s  ] %s. Failure Reason => %v"
	output := fmt.Sprintf(tmpl_msg, Red("\u2716"), Red("Failed"), check, Red(msg))
	log_output := fmt.Sprintf(tmpl_msg, "\u2716", "Failed", check, msg)
	LOGGER.Error(log_output)
	print(output)
}
