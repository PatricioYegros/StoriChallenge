package funcs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	reader := FileReader{}

	tests := []struct {
		setup      func(t *testing.T)
		name       string
		got        string
		want       [][]string
		err        error
		errMessage string
		reset      func(t *testing.T)
	}{
		{
			func(t *testing.T) {},
			"file not found",
			"not_found.csv",
			nil,
			ErrReadingFile,
			"error while reading file not_found.csv: open ../../csv/not_found.csv: El sistema no puede encontrar el archivo especificado",
			func(t *testing.T) {},
		},
		{
			func(t *testing.T) {
				writeFile(t, "test.txt", `asd,
asd`)
			},
			"file is not csv",
			"test.txt",
			nil,
			ErrReadingFile,
			"test.txt: record on line 2: wrong number of fields",
			func(t *testing.T) { removeFile(t, "test.txt") },
		},
		{
			func(t *testing.T) { writeFile(t, "short.csv", "Id,Date,Transaction") },
			"file not enough lines",
			"short.csv",
			nil,
			ErrReadingFile,
			"short.csv: file has less that 2 lines",
			func(t *testing.T) { removeFile(t, "short.csv") },
		},
		{
			func(t *testing.T) {
				writeFile(t, "testpass.csv", `Id,Date,Transaction
0,7/15,+60.5
1,7/28,-10.3
2,8/2,-20.46
3,8/13,+10`)
			},
			"correct file",
			"testpass.csv",
			[][]string{
				{"0", "7/15", "+60.5"},
				{"1", "7/28", "-10.3"},
				{"2", "8/2", "-20.46"},
				{"3", "8/13", "+10"},
			},
			nil,
			"",
			func(t *testing.T) { removeFile(t, "testpass.csv") },
		},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)

			ans, err := reader.Read(tt.got, "../../csv/")

			tt.reset(t)

			require.ErrorIs(t, err, tt.err)

			if err == nil {
				assert.ElementsMatch(t, ans, tt.want)
			} else {
				require.ErrorContains(t, err, tt.errMessage)
			}
		})
	}
}

func writeFile(t *testing.T, fileName string, content string) {
	file, err := os.Create(filepath.Join("../../csv/", filepath.Base(fileName)))
	require.NoError(t, err)
	_, err = file.WriteString(content)
	require.NoError(t, err)
	err = file.Close()
	require.NoError(t, err)
}

func removeFile(t *testing.T, fileName string) {
	err := os.Remove(filepath.Join("../../csv/", filepath.Base(fileName)))
	require.NoError(t, err)
}
