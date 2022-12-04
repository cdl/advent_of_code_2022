package utils

import "os"

func ReadFileString(filename string) string {
	input, err := os.ReadFile(filename)
	Check(err)

	return string(input)
}
