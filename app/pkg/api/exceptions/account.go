package exceptions

import "fmt"

// AccountAlreadyExistsException this exception describes that
// requested already exists and request cannot be performed
type AccountAlreadyExistsException struct{}

func (ex AccountAlreadyExistsException) Error() string {
	return fmt.Sprintf("Account already exists")
}

// VersionIncorrectException this exception describes that
// requested resource version is not correct
type VersionIncorrectException struct{}

func (ex VersionIncorrectException) Error() string {
	return fmt.Sprintf("Specified version incorrect")
}

// NotFoundException this exception describes that
// requested resource cannot be found or does not exist
type NotFoundException struct{}

func (ex NotFoundException) Error() string {
	return fmt.Sprintf("Specified resource does not exist")
}
