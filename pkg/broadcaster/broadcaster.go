package broadcaster

import (
	"fmt"
	"net/http"
)

func SayHi(url string) error {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return err
	}
	fmt.Printf("client: got responce!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	return nil
}
