// Package util contains utility functions
package util

import (
	"fmt"

	"github.com/gookit/color"
)

// PrintTabbed Prints a message prepended with a tab to STDOUT
func PrintTabbed(a ...any) {
	fmt.Printf("	%s\n", a...)
}

// PrintTabbedf Prints a message prepended with a tab to STDOUT when passed a format
func PrintTabbedf(f string, a ...any) {
	fmt.Printf("	%s", fmt.Sprintf(f, a...))
}

// PrintSuccess Prints a success message with a green indicator to STDOUT
func PrintSuccess(a ...any) {
	printIcon(color.FgGreen)
	fmt.Println(a...)
}

// PrintSuccessf Prints a success message with a green indicator to STDOUT when passed a format
func PrintSuccessf(f string, a ...any) {
	printIcon(color.FgGreen)
	fmt.Printf(f, a...)
}

// PrintInfo Prints an informational message with a blue indicator to STDOUT
func PrintInfo(a ...any) {
	printIcon(color.FgBlue)
	fmt.Println(a...)
}

// PrintInfof Prints an informational message with a blue indicator to STDOUT  when passed a format
func PrintInfof(f string, a ...any) {
	printIcon(color.FgBlue)
	fmt.Printf(f, a...)
}

// PrintWarning Prints a warning message with a yellow indicator to STDOUT
func PrintWarning(a ...any) {
	printIcon(color.FgYellow)
	fmt.Println(a...)
}

// PrintWarningf Prints a warning message with a yellow indicator to STDOUT when passed a format
func PrintWarningf(f string, a ...any) {
	printIcon(color.FgYellow)
	fmt.Printf(f, a...)
}

// PrintErr Prints an error message with a yellow indicator to STDOUT
func PrintErr(err error) {
	printIcon(color.FgRed)
	fmt.Println(err)
}

// PrintErrf Prints an error message with a yellow indicator to STDOUT when passed a format
func PrintErrf(f string, err error) {
	printIcon(color.FgRed)
	fmt.Printf(f, err)
}

// printIcon Prints a colored square icon to STDOUT
func printIcon(c color.Color) {
	color.Style{c}.Print("■ ")
}
