package internal

import "net/http"

func PostMessage(url string, message *Message) error {
	// Send the request
	_, err := http.Post(url, "application/json", nil)

	if err != nil {
		return err
	}

	return nil
}
