package content

import (
	"encoding/json"
	"net/http"

	"github.com/vexth1/handler"
)

type Content struct {
	Data string
}

func Hello(ctx *handler.Context) error {

	res := ctx.Response()
	res.WriteHeader(http.StatusOK)
	response := Content{"hello world! "}
	json, _ := json.Marshal(response)
	res.Write(json)

	return nil
}
