package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			// 处理错误
		}

		fmt.Println("Request Body:", string(body))

		// 对请求体做更多处理

		// 将处理结果写回响应
		_, err = writer.Write([]byte("Response Body"))
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		return
	}
}
