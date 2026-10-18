package string_array

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type StringArray []string

func (a *StringArray) Get() interface{} { return []string(*a) }

func (a *StringArray) Set(s string) error {
	*a = append(*a, s)
	return nil
}

func (a *StringArray) String() string {
	return strings.Join(*a, ",")
}

func (a *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	err := json.Unmarshal(bytes, a)
	return err
}

func (a *StringArray) Value() (driver.Value, error) {
	if a == nil || len(*a) == 0 {
		return nil, nil
	}
	return json.Marshal(*a)
}
