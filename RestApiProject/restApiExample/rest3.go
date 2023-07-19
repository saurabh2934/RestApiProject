package main

import (
	restful "github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

//This example shows how to have a program with 2 webServices containers
// each having a http server listening on its own port.
//
//The first "h                                                      ello" is added to the restful.DefaultContainer( and uses DefaultServeMux)
//For th e second "hello", a new container and ServeMux is created
//and requires a new http.Server with the container being the Handler.
//This first server is spawn in its own go-routine such that the program proceeds to create the second.
//
//GET http://localhost:8080/hello
//GET http://localhost:8081/hello

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/sk").To(hello))
	restful.Add(ws)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	container2 := restful.NewContainer()
	ws2 := new(restful.WebService)
	ws2.Route(ws2.GET("/s").To(hello2))

	container2.Add(ws2)
	server := &http.Server{Addr: ":8087", Handler: container2}
	log.Fatal(server.ListenAndServe())

	container3 := restful.NewContainer()
	Ws3 := new(restful.WebService)
	Ws3.Route(Ws3.GET("/s").To(hello3))
	container3.Add(Ws3)
	server2 := &http.Server{Addr: ":8088", Handler: container3}
	log.Fatal(server2.ListenAndServe())

}
func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "default world")
}

func hello2(req *restful.Request, resp *restful.Response) {
	resp.Write([]byte("please send the request"))

}
func hello3(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "third world")
}
