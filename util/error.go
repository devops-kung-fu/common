package util

import "log"

//IsErrorBool Checks to see if an error exists, and if so
//returns true after writing the error to the log
func IsErrorBool(err error) (b bool) {
	if err != nil {
		log.Println(err)
		b = true
	}
	return
}

//IfErrorLog Checks to see if an error exists, and if so
//simply writes it to the log.
func IfErrorLog(err error) {
	if err != nil {
		log.Println(err)
	}
}
