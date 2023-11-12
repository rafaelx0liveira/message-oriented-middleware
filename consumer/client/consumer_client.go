package client

import (
	"bytes"
	"net/http"
)

type ConsumerClient struct {
	HttpClient  *http.Client
	PostURL     string
	ContentType string
}

func NewConsumerClient(postURL string, contentType string) *ConsumerClient {
	return &ConsumerClient{
		HttpClient:  &http.Client{},
		PostURL:     postURL,
		ContentType: contentType,
	}
}

func (cc *ConsumerClient) PostWebhook(postBody []byte) error {
	request, err := http.NewRequest("POST", cc.PostURL, bytes.NewBuffer(postBody))

	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", cc.ContentType)

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
