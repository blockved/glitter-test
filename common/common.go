package common

import (
	"encoding/json"
	"fmt"
	"time"
)

func ConvToJSON(v interface{}) (str string) {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func PrintJson(formating string, val interface{}) {
	tmp, _ := json.Marshal(val)
	format := fmt.Sprintf("[%v]%v\n", time.Now(), formating)
	fmt.Printf(format, string(tmp))
}
