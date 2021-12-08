package helper

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	//log.Printf("JSON %v ", decoder)
	err := decoder.Decode(result)
	PanicIfError(err, "error helper/json at func readfromrequestbody when decode result")
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err, "error helper/json at func writetoresponsebody when encode response")
}
func ReadFromRequestQueryParam(request *http.Request, result interface{}) {
	var decoder = schema.NewDecoder()
	err := decoder.Decode(result, request.URL.Query())
	PanicIfError(err, "error helper/json at func readfromrequestqueryparam when decode url query")
}
