package storage

import (
	"fmt"
	"testing"
)

func TestKeyTable(t *testing.T) {

	tab, _ := NewKeyTab("./testDB")
	tab.Put("key", []string{"1", "2"})
	val := tab.Get("key")
	fmt.Println(val)

	raw := tab.GetAllRaw()
	fmt.Println(string(raw))

	tab.AppendBatch(map[string]string{
		"key": "3|4|1|5",
	})

	tab.Put("key1", []string{"5", "6", "7"})

	raw = tab.GetAllRaw()
	fmt.Println(string(raw))
}
