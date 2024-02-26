package broadcaster

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func sendRequest(jsonBody []byte, to string, path string, method string) error {
	bodyReader := bytes.NewReader(jsonBody)

	address := url.URL{Scheme: "http", Host: fmt.Sprintf("%s:3333", to), Path: path}

	req, err := http.NewRequest(method, address.String(), bodyReader)

	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)

	return err
}
