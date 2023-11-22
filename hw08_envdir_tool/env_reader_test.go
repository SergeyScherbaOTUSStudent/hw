package main

import (
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvReader(t *testing.T) {
	t.Run("not exist environment dir", func(t *testing.T) {
		input := path.Join("testdata", "nonexistent_file")
		actualEnv, e := ReadDir(input)

		require.Error(t, e)
		require.Equal(t, Environment(nil), actualEnv, "not equal")
	})

	testcases := []struct {
		title    string
		fileName string
		envValue EnvValue
	}{
		{"Test BAR file", "BAR", NewEnvValue("bar", true)},
		{"Test FOO file", "FOO", NewEnvValue("   foo\nwith new line", true)},
		{"Test HELLO file", "HELLO", NewEnvValue(`"hello"`, true)},
		{"Test UNSET file", "UNSET", NewEnvValue("", true)},
		{"Test EMPTY file", "EMPTY", NewEnvValue("", true)},
	}

	t.Run("Test ReadFileJustFirstLine", func(t *testing.T) {
		for _, testcase := range testcases {
			t.Run(testcase.title, func(t *testing.T) {
				from := path.Join("testdata", "env", testcase.fileName)
				res, err := ReadFileJustFirstLine(from)

				require.NoError(t, err)
				require.Equal(t, testcase.envValue, NewEnvValue(res, true), "EnvValue not equal")
			})
		}
	})

	t.Run("Test ReadDir", func(t *testing.T) {
		r := make(map[string]EnvValue)
		e := make(map[string]EnvValue)
		for _, testcase := range testcases {
			from := path.Join("testdata", "env", testcase.fileName)
			res, _ := ReadFileJustFirstLine(from)

			e[testcase.fileName] = testcase.envValue
			r[testcase.fileName] = NewEnvValue(res, true)
		}

		require.Equal(t, Environment(e), Environment(r), "environment not equal")
	})
}
