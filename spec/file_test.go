package spec_test

import (
	"fmt"
	"os"
)

func testReadFile(name string) []byte {
	data, err := os.ReadFile(name)

	if err != nil {
		panic(fmt.Sprintf("unable to read file %q: %s", name, err))
	}

	return data
}
