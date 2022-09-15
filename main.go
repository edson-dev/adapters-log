package adapters

import (
	"encoding/json"
	"go-pocketbase/pkg/adapters-log/log"
)

func main() {
	l := adapters.Log{}.New()
	s := `{"Name": "India", "test": "Test123"`
	v := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), v)
	l.Throw(err)
}
