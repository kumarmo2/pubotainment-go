package queueutils

import (
	queuemanager "pubwebservice/commonLibs/queue_manager"
	"sync"
)

var once sync.Once
var queueManager queuemanager.IQueueManager
var factory queueManagerFactory

type IQueueManagerFactory interface {
	GetQueueManager() queuemanager.IQueueManager
}

type queueManagerFactory struct{}

func initQueueManager() {
	queueManager = queuemanager.NewQueueManager()
	go queueManager.Start()
}

func (factory *queueManagerFactory) GetQueueManager() queuemanager.IQueueManager {
	once.Do(initQueueManager)
	return queueManager
}

func GetQueueManagerFactory() IQueueManagerFactory {
	return &factory
}
