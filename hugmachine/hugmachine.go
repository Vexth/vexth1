package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Vexth/vexth1/hugger"
	logging "github.com/op/go-logging"
)

func heartbreakerHandler(w http.ResponseWriter, req *http.Request) {
	logger.Infof("Meet heartbreaker from %s", req.Host)
	io.WriteString(w, hugger.Hugger())
}

var LISTENING_PORT = 3000
var logger = logging.MustGetLogger("hugmachine.log")

func main() {
	logging.NewLogBackend(os.Stderr, "", 0)
	http.HandleFunc("/heartbreaker", heartbreakerHandler)
	logger.Infof("Listening on port %d", LISTENING_PORT)
	fmt.Println(strconv.Itoa(LISTENING_PORT))
	err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(LISTENING_PORT), nil)
	if err != nil {
		logger.Fatal("ListenAndServe: " + err.Error())
	}
}
