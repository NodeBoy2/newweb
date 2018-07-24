package main

import (
	"io"
	"net/http"
	"os/exec"
	"log"
)

func reLauch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello Dog!</h1>")
	reLauch()
}

func main() {
	http.HandleFunc("/api", firstPage)
	http.ListenAndServe(":5000", nil)
}
