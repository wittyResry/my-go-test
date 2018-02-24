package exercise

import "github.com/Go-zh/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func MainMyReader() {
	reader.Validate(MyReader{}) // go get github.com/Go-zh/tour
}
