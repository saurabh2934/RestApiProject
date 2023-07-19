package main

//import _ "github.com/emicklei/go-restful/v3"
import (
	restful "github.com/emicklei/go-restful/v3"
	//_ "github.com/emicklei/go-restful/v3"
)

type data struct {
	d map[string]interface{}
}

func (d1 data) WebService() *restful.WebService {
	ws := new(restful.WebService)

	ws.
		Path("/calculation").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/").To(d1.ForAllData).
		Doc("get all data").
		Writes(d1.d).
		Returns(200, "OK", d1.d))

	ws.Route(ws.GET("/add").To(d1.AddTheMap).
		Doc("add the map"))
		//param(ws.PathParameter("key", "identifier of the key").DataType("string").DefaultValue("nil")).
		//Writes(d1.d).
		//Returns(200, "OK", d1.d).
		//Returns(404, "Not Found"))

	return ws
}
func (d1 data) AddTheMap(req *restful.Request, res *restful.Response) {
	//key := new(d1.d)

}
func (d1 data) ForAllData(req *restful.Request, res *restful.Response) {
	res.WriteEntity(d1)
}
func main() {

}
