// +build !android

package toast

import "fmt"

type t struct{}

func NewToaster() Toaster {
	return &t{}
}

func (t *t) Show(text string) error {
	// not implemented
	fmt.Println(text)
	return nil
}
