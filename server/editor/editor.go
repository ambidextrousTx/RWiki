package main

import (
	"net/http"
)

func Render() http.Handler {
	return http.FileServer(http.Dir("../../ui/edit.html"))
}

func main() {
	// placeholder

}
