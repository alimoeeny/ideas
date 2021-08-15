package ideas

import (
	"encoding/json"
	"fmt"
	"time"
)

type Factsheet interface {
	ID() string
	CurrentValue(key string) interface{}
	Reset() error
	String() string // for debugging
	CurrentVersion() int64
	VersionBump() int64
}

func NewDictionaryFactsheet(id string) *DictionaryFactsheet {
	return &DictionaryFactsheet{
		id:      id,
		version: int64(0),
		dic:     make(map[string]interface{}),
	}
}

type DictionaryFactsheet struct {
	id      string
	version int64
	dic     map[string]interface{}
}

func (dfs *DictionaryFactsheet) ID() string {
	return dfs.id
}

func (dfs *DictionaryFactsheet) CurrentValue(key string) interface{} {
	if dfs == nil || dfs.dic == nil {
		fmt.Println("This factsheet is not initialized properly")
		return nil
	}
	return dfs.dic[key]
}

func (dfs *DictionaryFactsheet) Reset() error {
	return nil
}

func (dfs *DictionaryFactsheet) SetValue(key string, value interface{}) {
	if dfs == nil || dfs.dic == nil {
		fmt.Println("This factsheet is not initialized properly, starting a new one with a new id")
		dfs = NewDictionaryFactsheet(newStrID())
	}
	dfs.dic[key] = value
}

func (dfs *DictionaryFactsheet) String() string {
	if dfs == nil || dfs.dic == nil {
		return "This factsheet is not initialized properly"
	}
	jDic := dfs.dic
	if jDic == nil {
		return "uninitialized factsheet"
	}
	jDic["id"] = dfs.id
	jBytes, err := json.Marshal(jDic)
	if err != nil {
		return err.Error()
	}
	return string(jBytes)
}

func (dfs *DictionaryFactsheet) CurrentVersion() int64 {
	if dfs == nil || dfs.dic == nil {
		return -1
	}
	return dfs.version
}

func (dfs *DictionaryFactsheet) VersionBump() int64 {
	if dfs == nil || dfs.dic == nil {
		return -1
	}
	dfs.version++
	return dfs.version
}

func (dfs *DictionaryFactsheet) MarshalJSON() ([]byte, error) {
	if dfs == nil || dfs.dic == nil {
		return nil, fmt.Errorf("this factsheet is not initialized properly")
	}
	tdic := dfs.dic
	tdic["_VERSION"] = dfs.version
	tdic["_ID"] = dfs.id
	tdic["_TIMESTAMP"] = time.Now().UnixNano()
	return json.Marshal(tdic)
}
