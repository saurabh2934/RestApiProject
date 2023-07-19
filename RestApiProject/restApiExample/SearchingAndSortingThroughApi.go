package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func setRootPath2(ws *restful.WebService, path string) {
	log.Println("setting root path:", path, " for webservice:", ws.Documentation())
	ws.Path(path)
}
func OperationOnSerAndSort(ws *restful.WebService, httpMethodName, apiPath, doc string, handler func(req *restful.Request, res *restful.Response)) {
	log.Println("Doing Operation: ", apiPath, "with http method:", httpMethodName)
	routeBuilder := ws.Doc(doc).Method(strings.ToUpper(httpMethodName)).Path(apiPath).To(handler)
	ws.Route(routeBuilder)
}
func WebService1() *restful.WebService {
	ws := new(restful.WebService)
	ws.Doc("Searching and Sorting")
	setRootPath2(ws, "/Ser&Sort")
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)
	OperationOnSerAndSort(ws, "POST", "/Searching", "Binary Search", binarySearch)
	OperationOnSerAndSort(ws, "POST", "/BubbleSorting", "Bubble Sort", BubbleSort)

	return ws
}
func binarySearch(req *restful.Request, res *restful.Response) {
	data := make(map[string]int)

	req.ReadEntity(&data)
	//dta := data["NO"]
	fmt.Println("Data Given by user", data)
	//making slice to store to the Data of User
	var dataSlice []int

	for _, er := range data {
		if er != 0 {
			dataSlice = append(dataSlice, int(er))
		}
	}

	fmt.Println("slice = ", dataSlice)

	str := ""

	for v, err := range dataSlice {
		if v != len(dataSlice)-1 {
			str = str + strconv.Itoa(err) + " "
		}
	}
	// doing binary search
	lenOfSlice := len(dataSlice)

	si := 0
	ei := lenOfSlice - 2
	find := lenOfSlice - 1
	for si <= ei || si > ei {
		if si > ei {
			res.WriteEntity("Element not Fount ")
			break
		} else {
			mid := (si + ei) / 2
			if dataSlice[find] == dataSlice[mid] {
				st := "Your Elements are [" + str + "] Element " + strconv.Itoa(dataSlice[find]) + " is available at index " + strconv.Itoa(mid)
				res.WriteEntity(st)
				break
			} else if dataSlice[find] > dataSlice[mid] {
				si = mid + 1
			} else {
				ei = mid - 1
			}
		}
	}

}

func BubbleSort(req *restful.Request, res *restful.Response) {
	data := make(map[string]int)
	req.ReadEntity(&data)
	var slicedata []int

	for _, er := range data {
		if er != 0 {
			slicedata = append(slicedata, er)
		}
	}
	fmt.Println(slicedata)
	for i := 0; i < len(slicedata)-1; i++ {
		for j := 0; j < len(slicedata)-i-1; j++ {
			if slicedata[j] > slicedata[j+1] {
				temp := slicedata[j]
				slicedata[j] = slicedata[j+1]
				slicedata[j+1] = temp
			}
		}
	}
	fmt.Println(slicedata)

	str := ""
	for v, err := range slicedata {
		if v != len(slicedata) {
			str = str + strconv.Itoa(err) + " "
		}
	}
	str1 := ""
	for _, err := range data {
		if err != 0 {
			str1 = str1 + strconv.Itoa(err) + " "
		}
	}

	res.WriteEntity(" unsorted array elements are: [" + str1 + "]")
	res.WriteEntity("\n sorted array elements are: [" + str + "]")

}

func selectionSort(req *restful.Request, res *restful.Response) {

}

func main() {
	fmt.Println("++++++++started main++++++++")
	restful.DefaultContainer.Add(WebService1())
	port := ":8082"

	log.Fatal(http.ListenAndServe(port, nil))

}
