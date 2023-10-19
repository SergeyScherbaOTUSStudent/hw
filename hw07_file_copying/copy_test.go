package main

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	from = path.Join("testdata", "input.txt")

	testcases := []struct {
		title     string
		from      string
		reference string
		offset    int64
		limit     int64
	}{
		{"Test_Offset0_limit0", from, "testdata/out_offset0_limit0.txt", 0, 0},
		{"Test_Offset0_limit10", from, "testdata/out_offset0_limit10.txt", 0, 10},
		{"Test_Offset0_limit1000", from, "testdata/out_offset0_limit1000.txt", 0, 1000},
		{"Test_Offset0_limit10000", from, "testdata/out_offset0_limit10000.txt", 0, 10000},
		{"Test_Offset100_limit1000", from, "testdata/out_offset100_limit1000.txt", 100, 1000},
		{"Test_Offset6000_limit1000", from, "testdata/out_offset6000_limit1000.txt", 6000, 1000},
	}

	for _, testcase := range testcases {
		to := path.Join("testdata", "out_tmp.txt")
		t.Run(testcase.title, func(t *testing.T) {
			err := Copy(testcase.from, to, testcase.offset, testcase.limit)
			defer os.Remove(to)

			a, _ := os.ReadFile(to)
			b, _ := os.ReadFile(testcase.reference)

			require.NoError(t, err)
			require.Equal(t, a, b)
		})
	}

	t.Run("file does not exist", func(t *testing.T) {
		err := Copy("testdata/input_.txt", "to", 0, 0)
		require.ErrorIs(t, err, ErrFileDoesNotExist)
	})

	t.Run("offset exceeds file size", func(t *testing.T) {
		err := Copy("testdata/input.txt", "to", 10000000, 0)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})
}
