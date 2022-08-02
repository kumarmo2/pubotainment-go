package queueutils

import (
	queuemanager "pubwebservice/commonLibs/queue_manager"
	"sync"
)

var once sync.Once
var factory queueManagerFactory

type IQueueManagerFactory interface {
	GetQueueManager(uniqueIdentifier string) queuemanager.IQueueManager
}

type queueManagerFactory struct {
	// TODO: this map should be thread safe.
	managers map[string]queuemanager.IQueueManager
}

func initFactory() {
	factory.managers = make(map[string]queuemanager.IQueueManager)
}

func (factory *queueManagerFactory) GetQueueManager(uniqueIdentifier string) queuemanager.IQueueManager {
	once.Do(initFactory)
	manager, ok := factory.managers[uniqueIdentifier] // ok tells whether the key is present in the map or not.
	if ok {
		return manager
	}

	manager = queuemanager.NewQueueManager()
	factory.managers[uniqueIdentifier] = manager
	go manager.Start()
	return manager
}

func GetQueueManagerFactory() IQueueManagerFactory {
	return &factory
}
