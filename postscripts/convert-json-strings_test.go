package postscripts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertJSONStrings(t *testing.T) {
	type Data struct {
		ID    int     `json:"id"`
		Word  string  `json:"word"`
		Code  bool    `json:"code"`
		Score float64 `json:"score"`
	}
	type Nest struct {
		JsonString Data `json:"json_string"`
	}
	type Resp struct {
		Ori         Data   `json:"ori"`
		JsonString  Data   `json:"json_string"`
		JsonStrings []Nest `json:"json_strings"`
	}

	var tcs = []struct {
		name     string
		input    map[string]interface{}
		expected Resp
	}{
		{
			name: "Positive test case",
			input: map[string]interface{}{
				"ori": map[string]interface{}{
					"id":    1,
					"word":  "a",
					"code":  false,
					"score": 0.1,
				},
				"json_string": "{\"id\":2,\"word\":\"b\",\"code\":true,\"score\":0.2}",
				"json_strings": []interface{}{
					map[string]interface{}{
						"json_string": "{\"id\":3,\"word\":\"c\",\"code\":false,\"score\":0.3}",
					},
					map[string]interface{}{
						"json_string": "{\"id\":4,\"word\":\"d\",\"code\":true,\"score\":0.4}",
					},
				},
			},
			expected: Resp{
				Ori: Data{
					ID:    1,
					Word:  "a",
					Code:  false,
					Score: 0.1,
				},
				JsonString: Data{
					ID:    2,
					Word:  "b",
					Code:  true,
					Score: 0.2,
				},
				JsonStrings: []Nest{
					{
						JsonString: Data{
							ID:    3,
							Word:  "c",
							Code:  false,
							Score: 0.3,
						},
					},
					{
						JsonString: Data{
							ID:    4,
							Word:  "d",
							Code:  true,
							Score: 0.4,
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			data := tc.input
			convertJSONStrings(data)

			b, err := json.Marshal(data)
			assert.NoError(t, err)

			actual := Resp{}
			err = json.Unmarshal(b, &actual)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
