package postscripts

import "encoding/json"

/*
convertJSONStrings converts JSON strings into JSON recursively.

from:

{"field": "{\"id\":17,\"hello\":\"word\"}"}

to:

{
	"field": {
		"hello": "word",
		"id": 17
	}
}
*/
func convertJSONStrings(data interface{}) {
	switch field := data.(type) {
	case map[string]interface{}:
		for k, v := range field {
			if s, ok := v.(string); ok {
				var temp interface{}
				if err := json.Unmarshal([]byte(s), &temp); err == nil {
					field[k] = temp
					convertJSONStrings(temp)
				}
			} else {
				convertJSONStrings(v)
			}
		}
	case []interface{}:
		for _, v := range field {
			convertJSONStrings(v)
		}
	}
}
