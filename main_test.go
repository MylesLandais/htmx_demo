package main

import (
	"net/http"
	"testing"
)

var num = 1000

func BenchmarkHttp(b *testing.B) {
	for i := 0; i < num; i++ {
		resp, err := http.Get("http://127.0.0.1:8080")
		if err != nil {
			return
		}
		defer resp.Body.Close()
	}

}
