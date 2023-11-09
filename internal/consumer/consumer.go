package consumer

import (
	"bytes"
	"container/list"
	"encoding/json"
	"net/http"
	"time"
)

type Consumer struct {
	LocalURL           string
	RemoteURL          string
	LaunchedHooks      list.List
	DefaultHookTimeOut int
}

func NewConsumer(localURL string, remoteURL string, timeOut int) *Consumer {
	return &Consumer{
		LocalURL:           localURL,
		RemoteURL:          remoteURL,
		LaunchedHooks:      *list.New(),
		DefaultHookTimeOut: timeOut,
	}
}

// Internal:
func (c *Consumer) processIncomingMsg(msg string) {
	print("Mensagem Recebida: " + msg)
}

// Public:
func (c *Consumer) HookCB(webhookResult *WebhookData) {
	listItem := c.LaunchedHooks.Front()

	for i := 0; i < c.LaunchedHooks.Len(); i++ {
		event := (listItem.Value).(*WebhookEvent)

		if event.PostData.EventID == webhookResult.EventID &&
			event.PostData.EventType == webhookResult.EventType {

			startTime := event.StartTime

			if time.Since(startTime) < time.Duration(event.Timeout)*time.Second {
				c.processIncomingMsg(webhookResult.MessageData)
				c.LaunchedHooks.Remove(listItem)
				break
			}
		}

		listItem = listItem.Next()
	}
}

func (c *Consumer) LaunchHook() error {
	data := WebhookData{
		EventID:     c.LaunchedHooks.Len(),
		EventType:   "consume_msg",
		MessageData: c.LocalURL,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(c.RemoteURL, "application/hook", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	webhookEvent := &WebhookEvent{
		PostData:  &data,
		Timeout:   c.DefaultHookTimeOut,
		StartTime: time.Now(),
	}
	c.LaunchedHooks.PushBack(webhookEvent)

	return err
}
