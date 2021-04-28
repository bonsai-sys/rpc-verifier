package errors

import "errors"

func CommunicationError() error {
	return errors.New("unable to reach distant server to authenticate user")
}

func Unauthorized() string {
	return "unauthorized"
}

func BadFormat() string {
	return "bad format"
}

func TokenExpired() string {
	return "token expired"
}

func InsufficentPermissions() string {
	return "insuffiencent permissions"
}

func WrongIssuer() string {
	return "wrong issuer"
}

func ServerError() string {
	return "unable to process verification"
}
