package kit

import (
	"fmt"

	"github.com/xuender/go-utils"
)

func ExampleNewDB() {
	tmpDir := fmt.Sprintf("/tmp/%s", utils.NewID('D'))
	db, _ := NewDB(tmpDir)
	defer db.Close()

	key := []byte{'1'}

	fmt.Println(db.Has(key))
	db.Put(key, "txt")
	fmt.Println(db.Has(key))

	var str string
	db.Get(key, &str)
	fmt.Println(str)

	// Output:
	// false <nil>
	// true <nil>
	// txt
}
