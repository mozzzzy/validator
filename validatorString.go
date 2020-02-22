package validator

/*
 * Module Dependencies
 */

import (
	"errors"
	"fmt"
	"os"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

/*
 * Functions
 */

func StrLongerThan(str string, min int) error {
	if len(str) < min {
		return errors.New(fmt.Sprintf("Length of \"%v\" is shorter than %v.", str, min))
	}
	return nil
}

func StrShorterThan(str string, max int) error {
	if len(str) > max {
		return errors.New(fmt.Sprintf("Length of \"%v\" is longer than %v.", str, max))
	}
	return nil
}

func StrOneOf(str string, candidates []string) error {
	for _, candidate := range candidates {
		if candidate == str {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("\"%v\" is not found in %v.", str, candidates))
}

func FileExist(path string) error {
	_, err := os.Stat(path)
	return err
}

func FileNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}
	return errors.New(fmt.Sprintf("%v exists.", path))
}
