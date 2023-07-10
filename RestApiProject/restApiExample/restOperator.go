package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Operation(ws *restful.WebService, httpMethodName, apiPath, doc string, handler func(request *restful.Request, response *restful.Response)) {
	log.Println("Registering api:", apiPath, "with http method:", httpMethodName)
	routeBuilder := ws.Doc(doc).Method(strings.ToUpper(httpMethodName)).Path(apiPath).To(handler)
	ws.Route(routeBuilder)
}
func setRootPath1(ws *restful.WebService, path string) {
	log.Println("setting root path:", path, " for webservice:", ws.Documentation())
	ws.Path(path)
}

func WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Doc("Operator")
	setRootPath1(ws, "/op")
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)
	Operation(ws, "POST", "/add", "for adding data", Add)
	Operation(ws, "POST", "/sub", "for adding data", Subtract)

	return ws
}
func Divide(req *restful.Request, res *restful.Request) {
	//Division only on two number
	//data := make(map[string]int)
}

func Add(req *restful.Request, res *restful.Response) {
	data := make(map[string]int)
	err := req.ReadEntity(&data)
	fmt.Println(data)
	sum := 0
	for _, er := range data {
		if data != nil {
			sum = sum + er
		} else {
			break
		}
	}
	s := strconv.Itoa(sum)
	s1 := "Add = " + s
	if err == nil {
		res.WriteEntity(s1)
	}
}
func Subtract(req *restful.Request, res *restful.Response) {
	data := make(map[string]int)
	err := req.ReadEntity(&data)
	fmt.Println(data)
	sum := 0
	for _, er := range data {
		if data != nil {
			if sum == 0 {
				sum = er
			} else {
				sum = sum - er
			}
		} else {
			break
		}
	}
	s := strconv.Itoa(sum)
	s1 := "Subtract = " + s
	if err == nil {
		res.WriteEntity(s1)
	}
}
func main() {
	//log.Printf("/n ######")
	restful.DefaultContainer.Add(WebService())
	port := ":8089"

	log.Printf("start listining on#{port}")
	log.Fatal(http.ListenAndServe(port, nil))
}
