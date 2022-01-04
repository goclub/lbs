package main

import (
	"io"
	"log"
	"net/http"
)

var httpClient = &http.Client{}
func main () {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/favicon.ico" {
			writer.WriteHeader(500)
			return
		}

		url := "https://apis.map.qq.com"+ request.URL.Path + "?" + request.URL.Query().Encode()
		log.Print(url)
		proxyReq, err := http.NewRequest(request.Method, url, nil)  ; if err != nil {
		    return
		}
		resp, err := httpClient.Do(proxyReq) ; if err != nil {
			log.Print(err)
			writer.WriteHeader(500)
			return
		}
		writer.WriteHeader(resp.StatusCode)
		data, err := io.ReadAll(resp.Body) ; if err != nil {
		    log.Print(err)
		    writer.WriteHeader(500)
		    return
		}
		_, err = writer.Write(data) ; if err != nil {
			log.Print(err)
			writer.WriteHeader(500)
			return
		}
	})
	addr := ":8251"
	log.Print("http://127.0.0.1" + addr)
	log.Print(http.ListenAndServe(addr, nil))
}