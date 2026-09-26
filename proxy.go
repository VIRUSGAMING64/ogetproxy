package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func ServerHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.Header.Get("URI")
		scheme := r.Header.Get("scheme")
		if uri == "" {
			w.Write([]byte("{\"error\":\"invalid url\"}"))
			return
		}
		sch := scheme
		if sch == "" {
			sch = "https"
		}
		uri = sch + "://" + uri

		req, err := http.NewRequest(r.Method, uri, r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := r.Header.Get("headers")
		if data != "" {
			var headers map[string]string
			if err := json.Unmarshal([]byte(data), &headers); err != nil {
				http.Error(w, fmt.Sprintf("invalid headers: %v", err), http.StatusBadRequest)
				return
			}
			for key, value := range headers {
				req.Header.Set(key, value)
			}
		}

		OnRequest(w, req)

	})
}

func OnRequest(w http.ResponseWriter, req *http.Request) {
	cli := http.Client{
		Timeout: time.Second * 60,
	}
	res, err := cli.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(res.StatusCode)
	if _, err := io.Copy(w, res.Body); err != nil {
		fmt.Println(err)
	}
}

func main() {
	PORT := os.Getenv("PAGE_PORT")
	if PORT == "" {
		PORT = "8028"
	}
	server := http.Server{Addr: fmt.Sprintf(":%s", PORT), Handler: ServerHandler()}
	server.ListenAndServe()

}

/*
only for sergio22

*/
