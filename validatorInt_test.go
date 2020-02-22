package validator

/*
 * Module Dependencies
 */

import (
	"testing"

	"github.com/mozzzzy/testUtil"
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

func TestIntBiggerThan(t *testing.T) {
	t.Run("No error when check 3 >= 2", func(t *testing.T) {
		err := IntBiggerThan(3, 2)
		testUtil.NoError(t, err)
	})

	t.Run("No error when check 3 >= 3", func(t *testing.T) {
		err := IntBiggerThan(3, 3)
		testUtil.NoError(t, err)
	})

	t.Run("Error when check 3 >= 4", func(t *testing.T) {
		err := IntBiggerThan(3, 4)
		testUtil.WithError(t, err)
	})
}

func TestIntSmallerThan(t *testing.T) {
	t.Run("No error when check 3 >= 2", func(t *testing.T) {
		err := IntSmallerThan(2, 3)
		testUtil.NoError(t, err)
	})

	t.Run("No error when check 3 >= 3", func(t *testing.T) {
		err := IntSmallerThan(3, 3)
		testUtil.NoError(t, err)
	})

	t.Run("Error when check 3 >= 4", func(t *testing.T) {
		err := IntSmallerThan(4, 3)
		testUtil.WithError(t, err)
	})
}

func TestIntWithin(t *testing.T) {
	t.Run("Error when check 3 >= 1 >= 2", func(t *testing.T) {
		err := IntWithin(1, 2, 3)
		testUtil.WithError(t, err)
	})

	t.Run("No error when check 3 >= 2 >= 2", func(t *testing.T) {
		err := IntWithin(2, 2, 3)
		testUtil.NoError(t, err)
	})

	t.Run("No error when check 3 >= 3 >= 2", func(t *testing.T) {
		err := IntWithin(3, 2, 3)
		testUtil.NoError(t, err)
	})

	t.Run("Error when check 3 >= 4 >= 2", func(t *testing.T) {
		err := IntWithin(4, 2, 3)
		testUtil.WithError(t, err)
	})
}

func TestValidPort(t *testing.T) {
	t.Run("Error when check 0 is valid port number", func(t *testing.T) {
		err := IntValidPort(0)
		testUtil.WithError(t, err)
	})

	t.Run("No error when check 1 is valid port number", func(t *testing.T) {
		err := IntValidPort(1)
		testUtil.NoError(t, err)
	})

	t.Run("No error when check 65536 is valid port number", func(t *testing.T) {
		err := IntValidPort(65535)
		testUtil.NoError(t, err)
	})

	t.Run("Error when check 65536 is valid port number", func(t *testing.T) {
		err := IntValidPort(65536)
		testUtil.WithError(t, err)
	})
}
