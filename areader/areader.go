// Package areader implements io.Reader with a type that emits an infinite
// stream of the ASCII character 'A'. Adapted from tour.golang.org/methods/22.
//
// Level: beginner
// Topics: interfaces, io.Reader
package areader

import "fmt"

type B struct{}

func (B) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
		fmt.Printf("%v ", i)
	}
	return len(p), nil
}
