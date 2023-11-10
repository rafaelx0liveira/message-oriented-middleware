package util

import "fmt"

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("The %s param is required and must be a %s", name, typ)
}