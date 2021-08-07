package ideas

import (
	"encoding/json"
	"fmt"

	"cloud.google.com/go/datastore"
)

func (dfs *DictionaryFactsheet) Load(ps []datastore.Property) error {
	for _, p := range ps {
		switch p.Name {
		case "ID":
			dfs.id = p.Value.(string)
		case "DIC":
			dicJStr := p.Value.(string)
			err := json.Unmarshal([]byte(dicJStr), &dfs.dic)
			if err != nil {
				return err
			}
		default:
			msg := fmt.Sprintf("Property %s is not supported. Value: %v", p.Name, p.Value)
			fmt.Println(msg)
		}
	}
	return nil
}

func (dfs *DictionaryFactsheet) Save() ([]datastore.Property, error) {
	props := []datastore.Property{}
	props = append(props, datastore.Property{Name: "ID", Value: dfs.id})

	dicJStr, err := json.Marshal(dfs.dic)
	if err != nil {
		return props, err
	}

	props = append(props, datastore.Property{Name: "DIC", Value: dicJStr})
	return props, nil
}
