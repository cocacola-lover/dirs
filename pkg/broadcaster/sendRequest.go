package broadcaster

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func sendRequest(jsonBody []byte, to string, method string) error {
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(method, to, bodyReader)

	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return err
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return err
	}
	fmt.Printf("client: response body: %s\n", resBody)
	return nil
}