package main
import (
 "fmt"
 "log"
 "net/http"
 )
 func main() {
 fmt.Println("Hello,বাংলাদেশ")
 http.Handle("/", http.FileServer(http.Dir(".")))
 log.Printf("Starting httpweb Server on port 3005")
 log.Fatal(http.ListenAndServe(":3005", nil))
}
