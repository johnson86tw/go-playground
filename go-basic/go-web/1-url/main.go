package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/post", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	t := req.FormValue("token")
	io.WriteString(w, "This is your token:"+t)
}

func bar(w http.ResponseWriter, req *http.Request) {
	n := req.FormValue("name")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<body style="background:black; color: white"><form method="post">
        <input type="text" name="name">
        <input type="submit" value="confirm">
    </form></br>`+n+`</body>`)
}
