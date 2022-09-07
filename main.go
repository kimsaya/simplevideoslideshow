package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// configure the songs directory name and port
	const storagedisk = "/home/orangepi/videoapp/"
	const port = 8080

	// add a handler for the song files
	// http.Handle("/", addHeaders(http.FileServer(http.Dir("./"))))
	http.Handle("/", addHeaders(http.FileServer(http.Dir(storagedisk))))
	// http.Handle("/src/", addHeaders(http.FileServer(http.Dir(storagedisk))))
	
	fmt.Printf("Starting server on %v\n", port)
	// log.Printf("Serving %s on HTTP port: %v\n", storagedisk, port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

// func FileServer(c *http.Conn, req *http.Request) {
// 	if req.URL.Path == "/" {
// 		http.Redirect(c, "index.html", 307)
// 		return
// 	} else if strings.HasSuffix(req.URL.Path, "html") {
// 		data, _ := ioutil.ReadFile(req.URL.Path[1:])
// 		io.WriteString(c, string(data))
// 		return
// 	}
// 	http.ServeFile(c, req, req.URL.Path[1:])
// }
