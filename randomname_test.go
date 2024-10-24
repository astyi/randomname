package randomname

import (
	"fmt"
	"testing"
)

func TestGenerateName(t *testing.T) {
	for i := 0; i < 100; i++ {
		name := GenerateName()
		fmt.Println(name)

		name = GenerateNameNotExceedLen(8)
		fmt.Println(name)
	}
}
