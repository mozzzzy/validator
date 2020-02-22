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

func TestStrLongerThan(t *testing.T) {
	t.Run("No error when strlen of \"abcd\" >= 4", func(t *testing.T) {
		err := StrLongerThan("abcd", 4)
		testUtil.NoError(t, err)
	})
	t.Run("Error when strlen of \"abcd\" >= 5", func(t *testing.T) {
		err := StrLongerThan("abcd", 5)
		testUtil.WithError(t, err)
	})
}

func TestStrShorterThan(t *testing.T) {
	t.Run("No error when strlen of \"abcd\" <= 4", func(t *testing.T) {
		err := StrShorterThan("abcd", 4)
		testUtil.NoError(t, err)
	})
	t.Run("Error when strlen of \"abcd\" <= 3", func(t *testing.T) {
		err := StrShorterThan("abcd", 3)
		testUtil.WithError(t, err)
	})
}

func TestStrOneOf(t *testing.T) {
	t.Run("No error when \"ab\" is in {\"ab\", \"cd\", \"ef\"}", func(t *testing.T) {
		err := StrOneOf("ab", []string{"ab", "cd", "ef"})
		testUtil.NoError(t, err)
	})
	t.Run("Error when \"ab\" is in {\"cd\", \"ef\", \"gh\"}", func(t *testing.T) {
		err := StrOneOf("ab", []string{"cd", "ef", "gh"})
		testUtil.WithError(t, err)
	})
}

func TestFileExist(t *testing.T) {
	t.Run("No error when file exists", func(t *testing.T) {
		err := FileExist("./validatorString_test.go")
		testUtil.NoError(t, err)
	})

	t.Run("Error when file does not exist", func(t *testing.T) {
		err := FileExist("./not_found.go")
		testUtil.WithError(t, err)
	})
}

func TestFileNotExist(t *testing.T) {
	t.Run("No error when file does not exist", func(t *testing.T) {
		err := FileNotExist("./not_found.go")
		testUtil.NoError(t, err)
	})

	t.Run("Error when file exists", func(t *testing.T) {
		err := FileNotExist("./validatorString_test.go")
		testUtil.WithError(t, err)
	})
}
