# GO JSON POSTSCRIPT
Provide a list of postscripts for your JSON. You can also add / modify according to your needs.

## Prerequisite
`Go 1.21.7` or newer. You may be able to use different versions, but please watch out for compatibility.

## How to use
1. Determine your postscripts in `postscripts/init.go`
2. If you need a different custom, you may create a new one under `postscripts/` then do step (1). It's possible to execute multiple postscripts.
3. Run the service with `make run`
4. Provide the input in `files/input.json`
5. See the output in `files/output.json`. The output is automatically generated every time you save your input in `files/input.json`

## Example
### convertJSONStrings

Input
```
{
    "ori": {
        "id": 1,
        "word": "a",
        "code": false,
        "score": 0.1
    },
    "json_string": "{\"id\":2,\"word\":\"b\",\"code\":true,\"score\":0.2}",
    "json_strings": [
        {"json_string": "{\"id\":3,\"word\":\"c\",\"code\":false,\"score\":0.3}"},
        {"json_string": "{\"id\":4,\"word\":\"d\",\"code\":true,\"score\":0.4}"}
    ]
}
```

Output
```
{
	"json_string": {
		"code": true,
		"id": 2,
		"score": 0.2,
		"word": "b"
	},
	"json_strings": [
		{
			"json_string": {
				"code": false,
				"id": 3,
				"score": 0.3,
				"word": "c"
			}
		},
		{
			"json_string": {
				"code": true,
				"id": 4,
				"score": 0.4,
				"word": "d"
			}
		}
	],
	"ori": {
		"code": false,
		"id": 1,
		"score": 0.1,
		"word": "a"
	}
}
```