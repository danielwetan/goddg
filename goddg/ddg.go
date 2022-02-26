package goddg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var apiURL = "https://api.duckduckgo.com/?q=%s&format=json&pretty=1%s"

func (response *Response) Decode(body []byte) error {
	if err := json.Unmarshal(body, response); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func encodeURL(query string) string {
	queryEnc := url.QueryEscape(query)
	if strings.HasPrefix(query, "!") {
		return fmt.Sprintf(apiURL, queryEnc, "&no_redirect=1")
	}

	return fmt.Sprintf(apiURL, queryEnc, "")
}

func get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Query(query string) (*Response, error) {
	url := encodeURL(query)
	body, err := get(url)
	if err != nil {
		return nil, err
	}

	response := &Response{}
	if err = response.Decode(body); err != nil {
		return nil, err
	}

	return response, nil
}
