package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"github.com/dreamscached/minequery/v2"
	"time"
)

//export PingServer
func PingServer(host *C.char, port C.ushort, result **C.char, errMsg **C.char) C.int {
	goHost := C.GoString(host)
	goPort := uint16(port)

	pinger := minequery.NewPinger(
		minequery.WithTimeout(5*time.Second),
		minequery.WithUseStrict(true),
		minequery.WithProtocolVersion16(minequery.Ping16ProtocolVersion162),
		minequery.WithProtocolVersion17(minequery.Ping17ProtocolVersion172),
	)

	beta18, err := pinger.PingBeta18(goHost, int(goPort))
	if err != nil {
		*errMsg = C.CString(err.Error())
		return -1 // Error code
	}

	// Marshal the result into JSON
	resultJSON, err := json.Marshal(beta18)
	if err != nil {
		*errMsg = C.CString(err.Error())
		return -1
	}

	// Convert Go string to C string
	*result = C.CString(string(resultJSON))
	// Caller is responsible for freeing the returned strings
	return 0 // Success
}

func main() {}
