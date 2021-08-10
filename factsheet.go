package ideas

type Factsheet interface {
	ID() string
	CurrentValue(key string) interface{}
	Reset() error
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
		return nil
	}
	return dfs.dic[key]
}

func (dfs *DictionaryFactsheet) Reset() error {
	return nil
}

func (dfs *DictionaryFactsheet) SetValue(key string, value interface{}) {
	dfs.dic[key] = value
}
