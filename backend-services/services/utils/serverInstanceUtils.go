package utils

import (
	"sync"

	"github.com/google/uuid"
)

var once sync.Once
var serverInstanceId string

func GetServerId() string {
	once.Do(func() { serverInstanceId = uuid.New().String() })
	return serverInstanceId
}
