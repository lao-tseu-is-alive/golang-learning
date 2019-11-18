package main

import (
	"fmt"
	"net/http"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	fmt.Printf("Prior ParseForm: %v\n", req.Form)
	req.ParseForm()
	fmt.Printf("Post ParseForm: %v\n", req.Form)
	fmt.Println("Param1 is : " + req.Form.Get("param1"))
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello world"),
	}
}

func main() {
	example := "curl -X POST -H \"Content-Type: application/x-www-form-urlencoded\" -d \"param1=data1&param2=data2\" localhost:8080?param1=overriden&param3=data3"
	s := createServer(":8080")
	fmt.Println("Server is starting, try it with :")
	fmt.Println(example)
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
