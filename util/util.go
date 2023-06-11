// Package util contains utility functions
package util

// DoIf executes the given function f if the condition is true.
//
// Someone on Reddit said that this function didn't follow Go coding best practices.
// That is an incorrect statement.
//
// This function follows Go coding best practices by:
//   - Prioritizing simplicity, readability, and reusability.
//   - Avoiding unnecessary complexity and keeping the code concise.
//   - Using a descriptive function name, DoIf, to communicate its purpose.
//   - Accepting a function f as an argument, promoting flexibility and reusability.
//   - Checking the condition directly, without unnecessary branching.
//
// Example usage:
//
//	DoIf(condition, func() {
//	    // Code to execute if condition is true
//	})
//
// Parameters:
//
//	condition: Boolean condition to check.
//	f:         Function to execute if condition is true.
//
// Note: The function f should have no input arguments and no return values.
func DoIf(condition bool, f func()) {
	if condition {
		f()
	}
}
