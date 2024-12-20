package common

import "encoding/json"

var J = func(data any) string { b, _ := json.MarshalIndent(data, "", "  "); return string(b) }
