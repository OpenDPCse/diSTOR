package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"strings"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

type KeyMap map[string]string

type KeyTable struct {
	db *leveldb.DB
	mu sync.Mutex
}

func NewKeyTab(path string) (*KeyTable, error) {

	db, err := leveldb.OpenFile(path, &opt.Options{
		BlockSize:           2 * opt.MiB,
		CompactionTotalSize: 12 * opt.MiB,
	})
	if err != nil {
		return nil, err
	}

	return &KeyTable{
		db: db,
		mu: sync.Mutex{},
	}, nil
}

func (k *KeyTable) Put(key string, cids []string) bool {
	val := k.toBytes(cids)
	err := k.db.Put([]byte(key), val, nil)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (k *KeyTable) Get(key string) []string {
	value, err := k.db.Get([]byte(key), nil)
	if err != nil {
		return nil
	} else {
		return k.split(value)
	}
}

func (k *KeyTable) Append(key string, cids []string) bool {
	k.mu.Lock()
	defer k.mu.Unlock()
	val := k.Get(key)

	m := make(map[string]struct{}, 5)

	for _, s := range val {
		m[s] = struct{}{}
	}

	for _, cid := range cids {
		if _, ok := m[cid]; !ok {
			val = append(val, cid)
		}
	}
	return k.Put(key, val)
}

func (k *KeyTable) AppendBatch(km KeyMap) {
	for key, val := range km {
		k.Append(key, k.split([]byte(val)))
	}
}

func (k *KeyTable) AppendBatchRaw(raw []byte) {
	km := make(KeyMap, 10)
	json.Unmarshal(raw, &km)
	k.AppendBatch(km)
}

func (k *KeyTable) GetAll() KeyMap {

	km := make(map[string]string, 10)
	iterator := k.db.NewIterator(nil, nil)
	for iterator.Next() {
		km[string(iterator.Key())] = string(iterator.Value())
	}

	return km
}

func (k *KeyTable) GetAllRaw() []byte {
	km := k.GetAll()
	raw, _ := json.Marshal(km)
	return raw
}

func (k *KeyTable) GetAllKeys() []string {
	km := k.GetAll()
	keys := make([]string, 0, 15)
	for key, _ := range km {
		keys = append(keys, key)
	}
	return keys
}

func (k *KeyTable) split(val []byte) []string {
	if len(val) == 0 {
		return nil
	}
	s := strings.Split(string(val), "|")
	return s
}

func (k *KeyTable) toBytes(cids []string) []byte {
	var builder bytes.Buffer
	for _, v := range cids {
		builder.WriteString(v)
		builder.WriteByte('|')
	}
	return builder.Bytes()[:builder.Len()-1]
}

func (k *KeyTable) Close() {
	err := k.db.Close()
	if err != nil {
		fmt.Println(err)
	}
}
