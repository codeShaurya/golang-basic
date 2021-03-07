package main
import (
	"fmt"
	//"example.com/user/golang_basic/morestrings"
	//"github.com/google/go-cmp/cmp"
	"net/http"
)

//func main() {
//	i := 0
//	for i := 0; i < 4; i++ {
//		defer fmt.Print(i)
//	}
//	i++
//	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
//	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
//}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>welcome to awesome</h1>")
	fmt.Println(r)
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}



