package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func PrintAllElemernt(req *restful.Request, res *restful.Response) {
	type CustomArray map[string]interface{}
	str := ""
	mapArray := []CustomArray{}
	req.ReadEntity(&mapArray)
	fmt.Println(&mapArray)
	for _, mp := range mapArray {
		for k, v := range mp {
			switch c := v.(type) { // The special syntax switch c := v.(type) tells us that this is a type switch,
			// meaning that Go will try to match the type of v to each case in the switch statement.
			case string:
				v2 := fmt.Sprintf("%v", c) // converting interface to string
				str = str + "#" + k + ":" + v2
			case int:
				v2 := fmt.Sprintf("%v", c) // converting interface to string
				str = str + "#" + k + ":" + v2
			case float64:
				v2 := fmt.Sprintf("%v", c) // converting interface to string
				str = str + "#" + k + ":" + v2
			default:
				v2 := fmt.Sprintf("%v", c) // converting interface to string
				str = str + "#" + k + ":" + v2
			}
		}
	}
	fmt.Println(str)
	res.WriteAsJson(str)
}
func main() {
	wg := new(restful.WebService)
	wg.Route(wg.GET("/PrintAllTyep").To(PrintAllElemernt))
	restful.Add(wg)
	log.Fatal(http.ListenAndServe(":8087", nil))
}
