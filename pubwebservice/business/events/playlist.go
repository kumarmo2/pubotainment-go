package events

import (
	"log"
	queuemanager "pubwebservice/commonLibs/queue_manager"
	"sync"
	"time"
)

var once sync.Once
var queueManager queuemanager.IQueueManager

func initQueueManager() {
	queueManager = queuemanager.NewQueueManager()
	go queueManager.Start()
}

func payloadHandler(payload interface{}) interface{} {
	log.Printf("payload received: %v", payload)
	return payload
}

// TODO: change the return type of this function.
func GetEvents() interface{} {
	once.Do(initQueueManager)
	timeOutChan := make(chan bool)
	payloadChan := make(chan interface{})

	subscriber := queuemanager.NewSubscriber(func(payload interface{}) {
		result := payloadHandler(payload)
		payloadChan <- result
	})
	queueManager.AddSubscriber(subscriber)

	defer queueManager.RemoveSubscriber(subscriber)

	go func() {
		time.Sleep(time.Second * 30)
		timeOutChan <- true
	}()

	select {
	case _ = <-timeOutChan:
		log.Println("timeout while waiting for message")
		return map[string]string{"event": "timeout"}
	case result := <-payloadChan:
		return result
	}

	// Either Timeout happens or subscriber receives a new payload.
}
