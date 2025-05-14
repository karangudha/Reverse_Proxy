package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8080", "HTTP network address") // Define the flag

func main() {
	flag.Parse()

	http.HandleFunc("/", home)
	http.ListenAndServe(*addr, nil) // Use the parsed address
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	tokens := []string{"token1", "token2", "token3", "is", "am"}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		time.Sleep(time.Millisecond * 420)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "typer.html")
}
