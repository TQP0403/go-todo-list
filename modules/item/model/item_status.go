package model

import (
	"encoding/json"
	"strings"
)

type ItemStatus uint8

const (
	Processing ItemStatus = iota
	Done       ItemStatus = iota
	Canceled   ItemStatus = iota
)

var typeToString = map[ItemStatus]string{
	Processing: "0",
	Done:       "1",
	Canceled:   "2",
}

func (s *ItemStatus) String() string {
	switch *s {
	case Processing:
		return "Processing"
	case Done:
		return "Done"
	case Canceled:
		return "Canceled"
	default:
		return ""
	}
}

func (s *ItemStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(typeToString[*s])
}

func (s *ItemStatus) UnmarshalJSON(data []byte) (err error) {
	str := strings.Trim(string(data), `"`)

	switch str {
	case "0":
		*s = Processing
	case "Processing":
		*s = Processing
	case "1":
		*s = Done
	case "Done":
		*s = Done
	case "2":
		*s = Canceled
	case "Canceled":
		*s = Canceled
	}

	return nil
}
