package queuemanager

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

type IQueueManager interface {
	AddSubscriber(subscriber ISubscriber) error
	RemoveSubscriber(subscriber ISubscriber) error
	BroadCast(payload interface{}) error
	Start()
}

type ISubscriber interface {
	recieveMessage(payload interface{})
	getUniqueId() string
}

type subscriber struct {
	handler func(interface{})
	id      string
}

func (subscriber *subscriber) recieveMessage(payload interface{}) {
	subscriber.handler(payload)
}

func (subscriber *subscriber) getUniqueId() string { return subscriber.id }

func NewSubscriber(handler func(interface{})) ISubscriber {
	id, _ := uuid.NewUUID()
	return &subscriber{
		handler: handler,
		id:      id.String(),
	}

}

type queueManager struct {
	messages    chan interface{}
	quitChan    chan bool
	subscribers map[string]ISubscriber // TODO: this needs to be thread-safe map.
}

func (queueManager *queueManager) AddSubscriber(subscriber ISubscriber) error {
	if subscriber == nil {
		return errors.New("subscriber cannot be nil")
	}
	queueManager.subscribers[subscriber.getUniqueId()] = subscriber
	return nil
}

func (queueManager *queueManager) RemoveSubscriber(subscriber ISubscriber) error {
	if subscriber == nil {
		return errors.New("subscriber cannot be nil")
	}
	delete(queueManager.subscribers, subscriber.getUniqueId())
	return nil
}

func (queueManager *queueManager) clear() {
	queueManager.messages = nil
	queueManager.quitChan = nil
	queueManager.subscribers = nil
}

func (manager queueManager) startConsuming() {
	for {
		select {
		case payload := <-manager.messages:
			for _, consumer := range manager.subscribers {
				go consumer.recieveMessage(payload)
			}
		case _ = <-manager.quitChan:
			log.Println("Qutting")
			manager.clear()
			return
		}
	}
}

func (queueManager *queueManager) Start() {
	queueManager.startConsuming()
}

func (queueManager *queueManager) BroadCast(payload interface{}) error {
	if payload == nil {
		return errors.New("payload cannot be nil")
	}
	go func() {
		queueManager.messages <- payload
	}()
	return nil
}

func NewQueueManager() IQueueManager {
	return &queueManager{
		messages:    make(chan interface{}, 10),
		quitChan:    make(chan bool),
		subscribers: map[string]ISubscriber{},
	}
}
