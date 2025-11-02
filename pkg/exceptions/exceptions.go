package exceptions

import (
	"log"
)

// HandleAnError => providing an error handler among entire project
func HandleAnError(msg string) {
	// should be updated and
	log.Println("error handler msg -> ", msg)
	// os.Exit(1)
	// do some here
}
