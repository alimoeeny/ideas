package ideas

import (
	"encoding/json"
	"fmt"
)

type Factsheet interface {
	ID() string
	CurrentValue(key string) interface{}
	Reset() error
	String() string // for debugging
}

func NewDictionaryFactsheet(id string) *DictionaryFactsheet {
	return &DictionaryFactsheet{
		id:  id,
		dic: make(map[string]interface{}),
	}
}

type DictionaryFactsheet struct {
	id  string
	dic map[string]interface{}
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
