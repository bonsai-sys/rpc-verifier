package verifier

import "errors"

func CommunicationError() error {
	return errors.New("unable to reach distant server to authenticate user")
}
