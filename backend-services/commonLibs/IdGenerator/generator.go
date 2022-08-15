package idgenerator

import (
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

type IIdGenerator interface {
	New() int64
}

type idGenerator struct {
	node *snowflake.Node
}

func (gen *idGenerator) New() int64 {
	if gen == nil {
		panic("generator is nill")
	}
	if gen.node == nil {
		panic("generator.node is nill")
	}
	return gen.node.Generate().Int64()
}

var generator *idGenerator
var once sync.Once

func initGenerator() {
	// NOTE: if multiple instances of the service called time.Now() at the same time.
	// this will cause issue. Take care of this.
	node, err := snowflake.NewNode(time.Now().Unix() % 1024)
	if err != nil {
		panic(err.Error())
	}

	generator = &idGenerator{
		node: node,
	}
}

func GetIdGenerator() IIdGenerator {
	once.Do(initGenerator)
	return generator
}
