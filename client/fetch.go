package client

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var (
	once      sync.Once
	netClient *http.Client
)

//
// create a singleton http client to ensure
// maximum reuse of connection
//
func newNetClient() *http.Client {
	once.Do(func() {
		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 2 * time.Second,
		}
		netClient = &http.Client{
			Timeout:   time.Second * 2,
			Transport: netTransport,
		}
	})
	return netClient
}

//
// method - http method to invoke (post/put/get etc.)
// header - map of headers to include in request
// body - reader for any content to supply as request body
//
func Fetch(method string, url string, header map[string]string, body io.Reader) ([]byte, error) {

	// //
	// // TODO: turn off in production
	// //
	// fmt.Printf("\nmethod:%v\nurl:%v\n,header:%+v\n\n", method, url, header)

	// Create request.
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// //
	// // TODO: turn off in production
	// //
	// reqDump, err := httputil.DumpRequestOut(req, true)
	// if err != nil {
	//  fmt.Println("req-dump error: ", err)
	// }
	// fmt.Printf("\noutbound request\n\n%s\n\n", reqDump)

	// Add any required headers.
	for key, value := range header {
		req.Header.Add(key, value)
	}

	// Perform the network call.
	res, err := newNetClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// //
	// // TODO: turn off in production
	// //
	// responseDump, err := httputil.DumpResponse(res, true)
	// if err != nil {
	//  fmt.Println("resp-dump error: ", err)
	// }
	// fmt.Printf("\nresponse:\n\n%s\n\n", responseDump)

	// If response from network call is not 200, return error.
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Network call failed with response: %d", res.StatusCode))
	}

	// return response payload as bytes
	respByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read Fetch response")
	}

	return respByte, nil
}
