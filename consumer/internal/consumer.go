package consumer

import (
	"consumer/client"
	"container/list"
	"encoding/json"
	"errors"
	"sync"
	"time"
)

type Consumer struct {
	Client             *client.ConsumerClient
	LocalURL           string
	RemoteURL          string
	ListMutex          *sync.Mutex
	LaunchedHooks      list.List
	DefaultHookTimeOut int
}

const MAX_HOOKS = 1000

func NewConsumer(localURL string, remoteURL string, postURL string, timeOut int) *Consumer {
	return &Consumer{
		Client:             client.NewConsumerClient(remoteURL, postURL),
		LocalURL:           localURL,
		RemoteURL:          remoteURL,
		ListMutex:          &sync.Mutex{},
		LaunchedHooks:      *list.New(),
		DefaultHookTimeOut: timeOut,
	}
}

// Thread A
// Internal:
func (c *Consumer) processIncomingMsg(msg string) {
	print("Mensagem Recebida: " + msg)
}

// Public:
func (c *Consumer) HookCB(webhookResult *WebhookData) {
	c.ListMutex.Lock()
	listItem := c.LaunchedHooks.Front()

	println("CB recebido")
	println(webhookResult.EventID)
	println(webhookResult.EventType)
	println(webhookResult.MessageData)

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
	c.ListMutex.Unlock()
}

func (c *Consumer) LaunchHook() error {
	if c.LaunchedHooks.Len() >= MAX_HOOKS {
		panic(errors.New("Cant launch any more hooks"))
	}

	data := WebhookData{
		EventID:     c.LaunchedHooks.Len(),
		EventType:   "consume_msg",
		MessageData: c.LocalURL,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	errPost := c.Client.PostWebhook(jsonData)

	if errPost != nil {
		panic(errPost)
	}

	webhookEvent := &WebhookEvent{
		PostData:  &data,
		Timeout:   c.DefaultHookTimeOut,
		StartTime: time.Now(),
	}
	c.LaunchedHooks.PushBack(webhookEvent)

	return err
}

// Thread B
func (c *Consumer) checkExpiredHooks() {

	listItem := c.LaunchedHooks.Front()

	for i := 0; i < c.LaunchedHooks.Len() && listItem != nil; i++ {
		event := (listItem.Value).(*WebhookEvent)

		startTime := event.StartTime

		if time.Since(startTime) < time.Duration(event.Timeout)*time.Second {
			c.LaunchedHooks.Remove(listItem)
			print("Removido Webhook expirado!")
		}

		listItem = listItem.Next()
	}
}

// launchInterval in ms
func (c *Consumer) LaunchHooksPeriodically(launchInterval int) {
	for {
		c.ListMutex.Lock()
		err := c.LaunchHook()
		c.ListMutex.Unlock()
		if err != nil {
			panic(err)
		}

		c.ListMutex.Lock()
		c.checkExpiredHooks()
		c.ListMutex.Unlock()

		time.Sleep(time.Millisecond * time.Duration(launchInterval))
	}
}
