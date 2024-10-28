package excepriton

import "log"

// HandleAnError => providing an error handler among entire project
func HandleAnError(msg string, err error) {
	// should be updated and
	log.Println(msg, err)
	// do some here
}
