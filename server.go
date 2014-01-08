package main
 
import (
  "os"
  "fmt"
  "net/http"
)

type Hello struct{}

// Implement the method ServeHTTP on struct Hello.
func (h Hello) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request) {
    fmt.Fprint(w, "Hello!\n")
}


func main() {
  port := os.Getenv("PORT")
  fmt.Println("PORT:", port)
  
  var h Hello
  http.ListenAndServe(":" + port, h)
}