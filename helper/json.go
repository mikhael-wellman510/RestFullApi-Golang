package helper

import (
	"encoding/json"
	"net/http"
)

// Read from request , dan ubah dari JSON ke tipe data yang di tuju
func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfErr(err)
}

// Ubah data nya ke JSON
func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfErr(err)
}
