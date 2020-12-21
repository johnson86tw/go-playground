package main

import (
	"context"
	"net/http"
)

type reqKey string

func main() {
	helloworldHandler := http.HandlerFunc(helloworld)
	http.Handle("/welcome", injectMsgID(helloworldHandler))
	http.ListenAndServe(":8080", nil)
}

func helloworld(res http.ResponseWriter, req *http.Request) {
	msgID := ""
	if msg := req.Context().Value(reqKey("msg-id")); msg != nil {
		if val, ok := msg.(string); ok {
			msgID = string(val)
		}
	}

	res.Header().Add("Msg-Id", msgID)
	res.Write([]byte("Hello World!\n"))
}

func injectMsgID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		msgID := "123"
		ctx := context.WithValue(req.Context(), reqKey("msg-id"), msgID)
		request := req.WithContext(ctx)
		next.ServeHTTP(res, request)
	})
}
