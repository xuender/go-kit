package kit

import (
	"fmt"

	face "github.com/Kagami/go-face"
)

func Example1() {
	db, _ := NewDB("db")
	defer db.Close()
	faces := []*Face{}
	db.Get([]byte("faces"), &faces)
	fmt.Println(len(faces))

	// Output:
	// false <nil>
}
func ExampleClassifier_Classify() {
	db, _ := NewDB("db")
	defer db.Close()
	faces := []*Face{}
	db.Get([]byte("faces"), &faces)

	rec, _ := face.NewRecognizer("data")
	defer rec.Close()
	c := Classifier{
		Rec:     rec,
		Named:   map[string][]Face{},
		Waiting: [][]Face{},
	}
	c.Classify(faces)

	// Output:
	// false <nil>
}

// func ExampleClassifier() {
// 	rec, _ := face.NewRecognizer("data")
// 	defer rec.Close()
// 	db, _ := NewDB("db")
// 	defer db.Close()
// 	suffix := "JPG"
// 	faces := []*Face{}
// 	err := filepath.Walk("26", func(filename string, fi os.FileInfo, err error) error {
// 		if fi.IsDir() {
// 			return nil
// 		}
// 		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
// 			fmt.Println(filename)
// 			p, err := os.Open(filename)
// 			if err != nil {
// 				return err
// 			}
// 			defer p.Close()
// 			m, _, err := image.Decode(p)
// 			if err != nil {
// 				return err
// 			}
// 			rgb := m.(*image.YCbCr)
// 			if fs, err := rec.RecognizeFile(filename); err == nil {
// 				for _, f := range fs {
// 					sub := rgb.SubImage(f.Rectangle).(*image.YCbCr)
// 					i := len(faces)
// 					sf, _ := os.Create(fmt.Sprintf("sub/%d.jpg", i))
// 					defer sf.Close()
// 					jpeg.Encode(sf, sub, &jpeg.Options{100})
// 					faces = append(faces, NewFace(f, []byte(strconv.Itoa(i))))
// 				}
// 			}
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	db.Put([]byte("faces"), faces)

// 	fmt.Println("xx")

// 	// Output:
// 	// false <nil>
// 	// true <nil>
// 	// txt
// }
