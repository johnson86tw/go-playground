package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// io.Copy
	http.Handle("/gin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("gin.jpg")
		if err != nil {
			fmt.Fprintln(w, "file upload fail")
		}

		io.Copy(w, f)
	}))

	// http.ServeContent
	http.Handle("/scientist", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("scientist.jpg")
		if err != nil {
			log.Fatalln(err)
		}

		fi, err := f.Stat()
		if err != nil {
			log.Fatalln(err)
		}

		http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
	}))

	// http.ServeFile
	http.Handle("/hawking", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "scientist.jpg")
	}))

	// http.FileServer
	// 不知道為什麼router一定要寫根目錄
	http.Handle("/", http.FileServer(http.Dir("./public")))

	// StripPrefix
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))

	// 更不知道為什麼瀏覽器搜尋/home一直跳/home/然後就是not found除非把router改成/home/
	// Chrome就可以...
	http.Handle("/home/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, "<img style='width: 300px' src='/gin' /><script src='./resources/main.js'></script>")
	}))

	http.ListenAndServe(":8080", nil)
}
