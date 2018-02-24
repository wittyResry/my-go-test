package readers

import (
	"fmt"
	"io"
	"strings"
)

func MainReader() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("%v %v %v\n", n, err, b)
			fmt.Printf("b[:n] = %q\n", b[:n])
		}
	}
}
