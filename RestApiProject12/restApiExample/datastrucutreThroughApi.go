package main

import (
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	RestApiProject "m2-plus.com/saurabh"
	"net/http"
	"strconv"
)

func subpathforWork(path, subPath string) string {
	return path + subPath
}
func OperationOnDS(wg *restful.WebService, httpMethod, apiPath string, doc string, handler func(req *restful.Request, res *restful.Response)) {
	wg.Route(wg.Method(httpMethod).Path(apiPath).Doc(doc).To(handler))

}
func setRootPath11(wg *restful.WebService, st string) {
	log.Println("setting root path:", wg, " for webservice:", wg.Documentation())
	wg.Path(st)
}

func WebService12() *restful.WebService {
	wg := new(restful.WebService)
	wg.Doc("Data Structure")
	setRootPath11(wg, "/DataStructure")
	wg.Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)
	OperationOnDS(wg, "POST", subpathforWork("/Searching", "/linear"), "ds serachig", NormalSearch)         // /Searching/normal
	OperationOnDS(wg, "POST", subpathforWork("/Searching", "/binarySearch"), "binary search", BinarySearch) // /Searching/binarySearch
	OperationOnDS(wg, "POST", subpathforWork("/Sorting", "/bubbleSort"), "bobble sort", bubbleSort)
	OperationOnDS(wg, "POST", subpathforWork("/Sorting", "/selectionSort"), "Selection sort", SelectionSort)
	OperationOnDS(wg, "POST", subpathforWork("/Sorting", "/insertionSort"), "insertion sort", insertionSort)
	OperationOnDS(wg, "POST", subpathforWork("/Sorting", "/quickSort"), "Quick Sort", QuickSort)
	OperationOnDS(wg, "POST", subpathforWork("/Sorting", "/mergeSort"), "Merge Sort", MergeSort)

	return wg
}

// http://localhost:8080/DataStrucure/Searching/linear
func NormalSearch(req *restful.Request, res *restful.Response) {
	type Array struct {
		Arr []int `json:"array"`
		F   int   `json:"find"`
	}
	ar := Array{}       // creating object for Array structure
	req.ReadEntity(&ar) // storing json data to structure
	fmt.Println(ar)
	find := ar.F                                       // initializing find variable which to be search
	er := RestApiProject.LinearSearching(ar.Arr, find) // calling function of linear search to search element
	log.Fatal(er)
	if er != nil {
		res.WriteAsJson(er.Error())
	}
}

// http://localhost:8089/DataStrucure/Searching/binarySearch
func BinarySearch(req *restful.Request, res *restful.Response) {
	type Array struct {
		Arr []int `json:"array"`
		F   int   `json:"find"`
	}
	ar := Array{}       // creating object for Array structure
	req.ReadEntity(&ar) // storing json data to structure
	fmt.Println(ar)
	find := ar.F // initializing find variable which to be search
	er := RestApiProject.BinarySearch(ar.Arr, find)
	fmt.Println(er)
	if er != nil {
		res.WriteAsJson(er.Error())
	}
}

// http://localhost:8080/DataStrucure/Sorting/bubbleSort
func bubbleSort(req *restful.Request, res *restful.Response) {
	type Array struct {
		Arr []int `json:"array"`
	}
	ar := Array{}
	req.ReadEntity(&ar)
	fmt.Println(ar)
	ar1 := RestApiProject.BubbleSorting(ar.Arr)
	fmt.Println(ar)
	res.WriteEntity(ar1)
}

func SelectionSort(req *restful.Request, res *restful.Response) {
	type Array struct {
		Arr []int `json:"array"`
	}
	ar := Array{}
	req.ReadEntity(&ar)
	fmt.Println(ar)
	RestApiProject.SelSort(ar.Arr) // Calling SelSort function to sort the array using Selection sorting technique
	res.WriteAsJson(ar)
}
func insertionSort(req *restful.Request, res *restful.Response) {
	type Array struct {
		Arr []int `json:"array"`
	}
	ar := Array{}
	req.ReadEntity(&ar) // request by server for (taking data send by server as JSON form)
	fmt.Println(ar)
	RestApiProject.InsertionSoting(ar.Arr)
	res.WriteAsJson(ar) // sending response to server(as XML form)
}

func QuickSort(req *restful.Request, res *restful.Response) {
	type D struct {
		Arr []int `json:"array"`
	}
	data := D{}
	req.ReadEntity(&data)
	low := 0
	high := len(data.Arr) - 1
	ar := RestApiProject.QuickFunction(data.Arr, low, high)
	// let's convert array in string for sending data in json with masage
	stt := ""
	for _, er := range ar {
		if er != 0 {
			et := strconv.Itoa(er)
			stt = stt + et + " "
		}
	}
	res.WriteAsJson("Sorted array Elements are [" + stt + "]")
}
func MergeSort(req *restful.Request, res *restful.Response) {
	data := make(map[string]interface{}) // declaring an empty map which store string as key and value is of type interface
	req.ReadEntity(&data)
	fmt.Println(data)
	var sl []int               // declaring an empty slice
	for _, dta := range data { // iterating over the map
		switch c := dta.(type) { // doing switch operation
		case []interface{}: // if data sent by user is interface array
			for _, er := range c { // iterating the interface [] and we are getting data as json.number
				if er != nil {
					i, _ := er.(json.Number).Int64() // type casting json number to int64
					sl = append(sl, int(i))          // appending data into slice to sort it
				}
			}
		default: // if we want to pass map from the server
			if dta != nil {
				i, _ := dta.(json.Number).Int64()
				sl = append(sl, int(i))
			}
		}
	}
	s := RestApiProject.MergerSort(sl)
	res.WriteEntity(s)
}

func main() {
	restful.DefaultContainer.Add(WebService12())
	log.Fatal(http.ListenAndServe(":8089", nil))
}
