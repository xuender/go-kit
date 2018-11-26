package kit

import (
	"fmt"

	"github.com/Kagami/go-face"
)

func ExampleNewPhoto() {
	rec, _ := face.NewRecognizer("data")
	defer rec.Close()

	b, err := NewPhoto("bona2.jpg", rec)
	fmt.Println(err)
	fmt.Printf("%d\n", len(b.Faces))

	p, err := NewPhoto("has-lens-info.jpg", rec)
	fmt.Println(err)
	fmt.Println(p.Ca)
	fmt.Println(p.Make, p.Model)

	// Output:
	// <nil>
	// 1
	// <nil>
	// 2014-09-01 15:03:47 +0800 CST
	// Apple iPhone 4S

}
