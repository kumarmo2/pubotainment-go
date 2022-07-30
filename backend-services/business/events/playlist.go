package events

import (
	"log"
	queueUtils "pubwebservice/business/queue_utils"
	queuemanager "pubwebservice/commonLibs/queue_manager"
	"time"
)

func payloadHandler(payload interface{}) interface{} {
	log.Printf("payload received: %v", payload)
	return payload
}

// TODO: change the return type of this function.
func GetEvents() []interface{} {

	factory := queueUtils.GetQueueManagerFactory()
	queueManager := factory.GetQueueManager()
	timeOutChan := make(chan bool)
	payloadChan := make(chan interface{})

	subscriber := queuemanager.NewSubscriber(func(payload interface{}) {
		result := payloadHandler(payload)
		payloadChan <- result
	})
	queueManager.AddSubscriber(subscriber)

	defer queueManager.RemoveSubscriber(subscriber)

	go func() {
		time.Sleep(time.Second * 5)
		timeOutChan <- true
	}()

	select {
	case _ = <-timeOutChan:
		log.Println("timeout while waiting for message")
		return []interface{}{map[string]string{"event": "timeout"}}
	case result := <-payloadChan:
		return []interface{}{map[string]interface{}{"event": result}}
	}

	// Either Timeout happens or subscriber receives a new payload.
}
