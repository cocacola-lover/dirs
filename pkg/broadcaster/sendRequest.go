package broadcaster

import (
	"bytes"
	"net/http"
)

func sendRequest(jsonBody []byte, to string, method string) error {
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(method, to, bodyReader)

	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)

	return err
}
