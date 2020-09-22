package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RestClient -
var (
	RestClient restInterface
)

// RestInterface defines a contract for REST clients
type restInterface interface {
	Post(URL string, body interface{}, header http.Header) (*http.Response, error)
}

func init() {
	RestClient = &restClient{}
}

type restClient struct{}

// Post makes POST requests
func (rc *restClient) Post(URL string, body interface{}, header http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, URL, bytes.NewReader(jsonBytes))
	request.Header = header

	client := http.Client{}
	return client.Do(request)
}
