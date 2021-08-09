package ideas

type Factsheet interface {
	ID() string
	CurrentValue(key string) interface{}
	Reset() error
}

func NewDictionaryFactsheet() *DictionaryFactsheet {
	return &DictionaryFactsheet{
		id:  newStrID(),
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
	return dfs.dic[key]
}

func (dfs *DictionaryFactsheet) Reset() error {
	return nil
}

func (dfs *DictionaryFactsheet) SetValue(key string, value interface{}) {
	dfs.dic[key] = value
}
