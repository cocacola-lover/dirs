package broadcaster

import (
	"bytes"
	drequests "dirs/pkg/requests"
	dtasks "dirs/pkg/tasks"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ProcessAskInfoTask(task *dtasks.AskInfoTask) error {
	jsonBody, err := json.Marshal(&drequests.SendInfoRequest{Search: task.Search, Info: *task.Result})
	if err != nil {
		fmt.Printf("client: could not marshal request: %s\n", err)
		return err
	}

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, task.From, bodyReader)

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
