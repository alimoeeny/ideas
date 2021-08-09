package ideas

import (
	"encoding/json"
	"fmt"

	"cloud.google.com/go/datastore"
)

// DictionaryFactsheet
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

// Measurment
func (m *Measurment) Load(ps []datastore.Property) error {
	for _, p := range ps {
		switch p.Name {
		case "ID":
			m.ID = p.Value.(string)
		case "timestamp":
			m.Timestamp = p.Value.(int64)
		case "DIC":
			dicJStr := p.Value.(string)
			var measurmentAsDic Measurment
			err := json.Unmarshal([]byte(dicJStr), &measurmentAsDic)
			if err != nil {
				return err
			}
			if m.ID != measurmentAsDic.ID {
				return fmt.Errorf("Something has gone horribly wrong here, measurment ID mismatch. Expected %s. Got %s", m.ID, measurmentAsDic.ID)
			}
			m.Value = measurmentAsDic.Value
			m.Unit = measurmentAsDic.Unit
		default:
			msg := fmt.Sprintf("Property %s is not supported. Value: %v", p.Name, p.Value)
			fmt.Println(msg)
		}
	}
	return nil
}

func (m *Measurment) Save() ([]datastore.Property, error) {
	props := []datastore.Property{}
	props = append(props, datastore.Property{Name: "ID", Value: m.ID})
	props = append(props, datastore.Property{Name: "timestamp", Value: m.Timestamp})

	dicJStr, err := json.Marshal(m)
	if err != nil {
		return props, err
	}

	props = append(props, datastore.Property{Name: "DIC", Value: dicJStr})
	return props, nil
}
