package utils

import "os"

func ReadFileString(filename string) string {
	input, err := os.ReadFile(filename)
	Check(err)

	return string(input)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
