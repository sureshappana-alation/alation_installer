/**
* This file contains code terminal prompts to get secret inputs.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	. "github.com/logrusorgru/aurora"
	"golang.org/x/term"
)

// This function loads the secret from environment variables if exist in normalized form
// environment variables normalized name has "_" instead of "." and is all uppercase
// example property name: alationanalytics.postgresPassword
// example expected environment variable: ALATIONANALYTICS_POSTGRESPASSWORD
// if not exists in environment variable function prompts the user to get the value
func getFromEnvOrPromptSecret(propertyName string, promptMsg string, required bool) string {
	inputValue := ""

	propertyValue, propertySetAsEnv := os.LookupEnv(normalizePropertyNameForEnvVariable(propertyName))

	if propertySetAsEnv {
		inputValue = propertyValue
	} else {
		fmt.Println(Yellow("=========================="))

		if required {
			inputValue = askSecret(promptMsg)
		} else {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println(Yellow(fmt.Sprintf("You can optionally enter the config value for %s", propertyName)))
			fmt.Print(Yellow("Enter \"yes\" or \"y\" to continue or \"no\" or anything else to skip: "))

			input, err := reader.ReadString('\n')

			if err != nil {
				panic(err)
			}

			input = strings.ToLower(strings.TrimSuffix(input, "\n"))

			if input == "yes" || input == "y" {
				inputValue = askSecret(promptMsg)
			} else {
				inputValue = ""
				fmt.Println("skipped!")
			}
		}
	}

	return inputValue
}

func askSecret(promptMsg string) string {
	userEnteredSecret1 := ""
	userEnteredSecret2 := ""
	fmt.Println("Input:", Yellow(promptMsg), Red("(Characters won't be displayed as you type)"))
	for {
		fmt.Print(Green("Please enter the value: "))
		userEnteredSecret1 = getSecret()
		fmt.Print(Green("Please confirm your value by entering it again: "))
		userEnteredSecret2 = getSecret()
		if userEnteredSecret1 != "" && userEnteredSecret1 == userEnteredSecret2 {
			break
		}
		fmt.Println(Red("Error: values are empty or are not matched. please try again."))
	}
	return userEnteredSecret1
}

func getSecret() string {
	// Get the initial state of the terminal.
	initialTermState, termErr := term.GetState(syscall.Stdin)
	if termErr != nil {
		panic(termErr)
	}

	// Restore it in the event of an interrupt.
	c := make(chan os.Signal)

	// The following line gives the lint issue "misuse of unbuffered os.Signal channel as argument to signal.Notify"
	// As this code is ported from AA team ignoring lint issue for now.
	//nolint
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		_ = term.Restore(syscall.Stdin, initialTermState)
		os.Exit(1)
	}()

	p, err := term.ReadPassword(syscall.Stdin)
	fmt.Println("")
	if err != nil {
		panic(err)
	}

	// Stop looking for ^C on the channel.
	signal.Stop(c)

	// Return the user entered value as a string.
	return string(p)
}

// this will replace "." with "_" and uppercase the whole string
func normalizePropertyNameForEnvVariable(propertyName string) string {
	return strings.ToUpper(strings.ReplaceAll(propertyName, ".", "_"))
}
