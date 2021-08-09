package ideas

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/datastore"
)

// DictionaryFactsheet
func (dfs *DictionaryFactsheet) Load(ps []datastore.Property) error {
	fmt.Printf("dfs-> %#v\n", dfs)
	for _, p := range ps {
		switch p.Name {
		case "ID":
			dfs.id = p.Value.(string)
		case "DIC":
			if dfs.dic == nil {
				dfs.dic = make(map[string]interface{})
			}
			fmt.Printf("-> %#v\n", p)
			fmt.Printf("-> %#v\n", p.Value)
			fmt.Printf("-> %t\n", p.Value)
			dicJBytes := p.Value.([]byte)
			dicJStr, err := base64.StdEncoding.DecodeString(string(dicJBytes))
			if err != nil {
				return err
			}
			err = json.Unmarshal(dicJStr, &dfs.dic)
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

	dicJBytes, err := json.Marshal(dfs.dic)
	if err != nil {
		return props, err
	}

	props = append(props, datastore.Property{Name: "DIC", Value: dicJBytes})
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
			dicJBytes := p.Value.([]byte)
			var measurmentAsDic Measurment
			err := json.Unmarshal([]byte(dicJBytes), &measurmentAsDic)
			if err != nil {
				return err
			}
			if m.ID != measurmentAsDic.ID {
				return fmt.Errorf("something has gone horribly wrong here, measurment ID mismatch. Expected %s. Got %s", m.ID, measurmentAsDic.ID)
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

	dicJBytes, err := json.Marshal(m)
	if err != nil {
		return props, err
	}

	props = append(props, datastore.Property{Name: "DIC", Value: dicJBytes})
	return props, nil
}
