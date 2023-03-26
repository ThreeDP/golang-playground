package main

import (
	"fmt"
	"io"
	"os"
	"syscall"
    "net/http"
)

func Greet(fd io.Writer, name string) {
	fmt.Fprintf(fd, "Hello, %s\n", name)
}

func GreetWithDup() {
	fd, _ := syscall.Open("ts.txt", syscall.O_RDWR, 755)
	syscall.Dup2(fd, 1)
	Greet(os.Stdout, "Julius")
}

func HandlerGreet(w http.ResponseWriter, r *http.Request) {
    Greet(w, "mundo")
}

func main() {
    err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerGreet))

    if err != nil {
        fmt.Println(err)
    }
}