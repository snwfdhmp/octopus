package resp

import (
	"io/ioutil"
	"net/http"
)

func Read(resp *http.Response) string {
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func Body(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	return Read(resp)
}
