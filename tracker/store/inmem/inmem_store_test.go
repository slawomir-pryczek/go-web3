package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"
	web3 "github.com/umbracle/go-web3"
	"github.com/umbracle/go-web3/tracker/store"
)

func TestInMemoryStore(t *testing.T) {
	store.TestStore(t, func(t *testing.T) (store.Store, func()) {
		return NewInmemStore(), func() {}
	})
}

func TestFlush(t *testing.T) {
	e := &Entry{
		logs: []*web3.Log{},
	}

	addLog := func(e *Entry, num uint64) {
		e.logs = append(e.logs, &web3.Log{BlockNumber: num})
	}

	// cannot flush if the array is empty
	assert.Len(t, e.Flush(0), 0)
	assert.Len(t, e.Flush(10), 0)

	addLog(e, 10)

	assert.Len(t, e.Flush(5), 0)
	assert.Len(t, e.Flush(15), 1)

	// the logs are empty now
	assert.Len(t, e.logs, 0)
}
