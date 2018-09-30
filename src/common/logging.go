package common

import (
	"fmt"
)

//Log the given message
func Log(message string) {
	//Simple wrapper over standard println but
	//could be expanded to, for example, write
	//to a log file
	fmt.Println(message)
}
