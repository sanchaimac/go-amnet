package fundconnext

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func (f *FundConnext) CallFcApiEOF(pass bool) error {
	fmt.Println("Func CallFcApiEOF")
	// Create a new mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Close the connection before sending any response
		conn, _, _ := w.(http.Hijacker).Hijack()
		conn.Close()
	}))
	defer mockServer.Close()

	// Make a request to the mock server
	client := &http.Client{}
	req, _ := http.NewRequest("GET", mockServer.URL, nil)
	// var res *http.Response
	res, err := client.Do(req)
	if err != nil {
		// Check if the error is EOF
		if strings.Contains(err.Error(), "EOF") {
			fmt.Println("EOF error occurred")
			retire := 3
			for i := 1; i <= retire; i++ {
				fmt.Println("retire count: ", i)
				res, err = client.Do(req)
				if i == 2 && pass {
					res = &http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(strings.NewReader("Test Response Body")),
					}
					err = nil
					break

				}
				time.Sleep(time.Duration(i) * time.Second)
			}
			if err != nil {
				fmt.Println("err: ", err)
				return errors.New(io.EOF.Error())
			}
		} else {
			fmt.Printf("Error: %s\n", err)
		}
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Response Body: %s\n", body)
	return nil
}
