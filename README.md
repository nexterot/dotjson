# Dotjson
Dotjson package provides DotFormat function for formatting json objects with "dot syntax"
# Usage
```
git clone https://github.com/nexterot/dotjson
cd dotjson
go build
```
Linux:
```
./dotjson -i input.json -o output.json
```
Windows:
```
dotjson -i input.json -o output.json
```
# Example
input.json
```json
{
    "a": {
        "b": {
            "c": 1,
            "d": 2
        }, 
        "c": "d",
        "f": "g"
    },
    "b": 100
}
```
output.json
```json
{
    "a.b.c": 1,
    "a.b.d": 2,
    "a.c": "d",
    "a.f": "g",
    "b": 100
}
```