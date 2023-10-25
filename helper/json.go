package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)

	// baca data json
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	// berithau bahwa ini json di header
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	// ubah ke json
	err := encoder.Encode(response)
	PanicIfError(err)
}
