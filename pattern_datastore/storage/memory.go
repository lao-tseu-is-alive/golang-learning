package storage

import (
	"encoding/json"
	"sync"
)

type inMemoryHashTable struct {
	m   map[int]*Todo
	seq int
	lck sync.RWMutex
}

func NewMemoryDB() (DB, error) {
	t := map[int]*Todo{}
	inMemTable := &inMemoryHashTable{m: t, seq: 0}
	return inMemTable, nil
}

func (i *inMemoryHashTable) New(val Todo) (string, error) {
	i.lck.Lock()
	defer i.lck.Unlock()
	i.seq++
	newTodo := Todo{
		ID:     i.seq,
		Title:  val.Title,
		IsDone: val.IsDone,
	}
	i.m[i.seq] = &newTodo
	res, _ := json.Marshal(i.m[i.seq])
	return string(res), nil
}

func (i *inMemoryHashTable) Get(key int) (string, error) {
	i.lck.RLock()
	defer i.lck.RUnlock()
	val, ok := i.m[key]
	if !ok {
		return "", ErrNotFound
	}
	res, _ := json.Marshal(val)
	return string(res), nil
}

func (i *inMemoryHashTable) List() (string, error) {
	i.lck.RLock()
	defer i.lck.RUnlock()
	res, _ := json.Marshal(i.m)
	return string(res), nil
}
