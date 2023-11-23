package hw07filecopying

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOffsetExceeds(t *testing.T) {
	to := "testdata/out_offset0_limit0.txt"
	err := Copy("testdata/input.txt", to, 123456, 10)
	require.EqualError(t, err, ErrOffsetExceedsFileSize.Error())
}

func TestCopy(t *testing.T) {
	tests := []struct {
		inputFile    string
		expectedFile string
		offset       int64
		limit        int64
	}{
		{
			inputFile:    "testdata/input.txt",
			expectedFile: "testdata/out_offset0_limit0.txt",
			offset:       0,
			limit:        0,
		},
		{
			inputFile:    "testdata/input.txt",
			expectedFile: "testdata/out_offset0_limit10.txt",
			offset:       0,
			limit:        10,
		},
		{
			inputFile:    "testdata/input.txt",
			expectedFile: "testdata/out_offset0_limit1000.txt",
			offset:       0,
			limit:        1000,
		},
		{
			inputFile:    "testdata/input.txt",
			expectedFile: "testdata/out_offset0_limit10000.txt",
			offset:       0,
			limit:        10000,
		},
		{
			inputFile:    "testdata/input.txt",
			expectedFile: "testdata/out_offset100_limit1000.txt",
			offset:       100,
			limit:        1000,
		},
		{
			inputFile:    "testdata/input.txt",
			expectedFile: "testdata/out_offset6000_limit1000.txt",
			offset:       6000,
			limit:        1000,
		},
	}

	//	HideProgress = true

	for i, test := range tests {
		t.Run("files should be equals", func(t *testing.T) {
			to := fmt.Sprintf("out%v_%v_%v.txt", i, test.limit, test.offset)
			defer os.Remove(to)

			Copy(test.inputFile, to, test.offset, test.limit)
			require.True(t, compare(test.expectedFile, to))
		})
	}
}

// Сравнение файлов - считаем хэш исходного и получившегося файла.
func compare(file1, file2 string) bool {
	hash1, err := getFileMD5(file1)
	if err != nil {
		return false
	}

	hash2, err := getFileMD5(file2)
	if err != nil {
		return false
	}

	return hash1 == hash2
}

func getFileMD5(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
