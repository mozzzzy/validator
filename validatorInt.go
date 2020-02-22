package validator

/*
 * Module Dependencies
 */

import (
	"errors"
	"fmt"
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

func IntBiggerThan(val, min int) error {
	if val < min {
		return errors.New(fmt.Sprintf("%v < %v.", val, min))
	}
	return nil
}

func IntSmallerThan(val, max int) error {
	if val > max {
		return errors.New(fmt.Sprintf("%v > %v.", val, max))
	}
	return nil
}

func IntWithin(val, min, max int) error {
	if BiggerErr := IntBiggerThan(val, min); BiggerErr != nil {
		return BiggerErr
	}
	if SmallerErr := IntSmallerThan(val, max); SmallerErr != nil {
		return SmallerErr
	}
	return nil
}

func IntValidPort(val int) error {
	return IntWithin(val, 1, 65535)
}
