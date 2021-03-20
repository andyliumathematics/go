package main

// import "net/http"
import "fmt"

func main3() {
	// http.ListenAndServe("127.0.0.1:80",handler);
	for i := 0; i < 1000; i++ {
		go fmt.Println(  i)
	}
}
