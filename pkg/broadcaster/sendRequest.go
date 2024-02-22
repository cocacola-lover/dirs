package broadcaster

import (
	"bytes"
	ss "dirs/pkg/serviceStore"
	"io"
	"net/http"
)

func sendRequest(jsonBody []byte, to string, method string, serviceStore ss.ServiceStore) error {
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(method, to, bodyReader)

	if err != nil {
		serviceStore.Logger.Error.Printf("client: could not create request: %s\n", err)
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		serviceStore.Logger.Error.Printf("client: error making http request: %s\n", err)
		return err
	}

	serviceStore.Logger.Info.Printf("client: got response!\n")
	serviceStore.Logger.Info.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		serviceStore.Logger.Error.Printf("client: could not read response body: %s\n", err)
		return err
	}
	serviceStore.Logger.Info.Printf("client: response body: %s\n", resBody)
	return nil
}
