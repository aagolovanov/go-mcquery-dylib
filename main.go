package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"context"
	"encoding/json"
	"github.com/mcstatus-io/mcutil/v4/status"
	"time"
)

//export GetPlayers
func GetPlayers(host *C.char, port C.ushort, result **C.char, errMsg **C.char) C.int {
	goHost := C.GoString(host)
	goPort := uint16(port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	response, err := status.Modern(ctx, goHost, goPort)
	if err != nil {
		*errMsg = C.CString(err.Error())
		return -1 // Error code
	}

	players := response.Players

	// Marshal the result into JSON
	res, err := json.Marshal(players)
	if err != nil {
		*errMsg = C.CString(err.Error())
		return -1
	}

	// Convert Go string to C string
	*result = C.CString(string(res))
	// Caller is responsible for freeing the returned strings
	return 0 // Success
}

func main() {}
