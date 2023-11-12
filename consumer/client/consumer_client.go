package client

import (
	"bytes"
	"net/http"
)

type ConsumerClient struct {
	HttpClient   *http.Client
	PostURL      string
	PostEndpoint string
}

func NewConsumerClient(postURL string, postEndpoint string) *ConsumerClient {
	return &ConsumerClient{
		HttpClient:   &http.Client{},
		PostURL:      postURL,
		PostEndpoint: postEndpoint,
	}
}

func (cc *ConsumerClient) PostWebhook(postBody []byte) error {
	request, err := http.NewRequest("POST", cc.PostURL, bytes.NewBuffer(postBody))

	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", cc.PostEndpoint)

	result, hcerr := cc.HttpClient.Do(request)

	if hcerr != nil {
		panic(hcerr)
	}

	defer result.Body.Close()

	if result.StatusCode != http.StatusCreated {
		panic(result.Status)
	}

	return nil
}
