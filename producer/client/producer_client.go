package client

import (
	"net/http"
	"bytes"
)

type ProducerClient struct {
	HttpClient   *http.Client
	PostURL      string
	PostEndpoint string
}

func NewProducerClient(postURL string, postEndpoint string) *ProducerClient {
	return &ProducerClient{
		HttpClient:   &http.Client{},
		PostURL:      postURL,
		PostEndpoint: postEndpoint,
	}
}

func (pr *ProducerClient) PostMsg(jsonMsg []byte) error {
	request, err := http.NewRequest("POST", pr.PostURL, bytes.NewBuffer(jsonMsg))

	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", pr.PostEndpoint)

	result, hcerr := pr.HttpClient.Do(request)

	if hcerr != nil {
		panic(hcerr)
	}

	defer result.Body.Close()

	if result.StatusCode != http.StatusCreated {
		panic(result.Status)
	}

	return nil
}