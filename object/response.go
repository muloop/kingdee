package object

import (
	"encoding/json"
	"fmt"
)

type Org struct {
	ID     int64  `json:"id"`     //金蝶组织内码
	Number string `json:"number"` //组织编码
	Name   string `json:"name"`   //组织名称
}

func (o *Org) UnmarshalJSON(data []byte) error {
	var raw []interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) != 3 {
		return fmt.Errorf("invalid array length: %d", len(raw))
	}

	o.ID, _ = raw[0].(int64)
	o.Number, _ = raw[1].(string)
	o.Name, _ = raw[2].(string)

	return nil
}
