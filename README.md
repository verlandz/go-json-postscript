# GO JSON POSTSCRIPT
Provide a list of postscripts for your JSON. You can also add / modify according to your needs.

## Prerequisite
`Go 1.21.7` or newer. You may be able to use different versions, but please watch out for compatibility.

## How to use
1. Provide the input in `files/input.json`
2. Determine your postscripts in `postscripts/init.go`
3. If you need a different custom, you may create a new one under `postscripts/` then do step (2). It's possible to execute multiple postscripts.
4. Run `make run`
5. See the output in `files/output.json`