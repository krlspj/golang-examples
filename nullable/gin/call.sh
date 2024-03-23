curl -X POST \
  http://localhost:8080/json \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "John Doe",
  "age": 0,
  "height": null,
  "married": false,
  "pet": "",
  "status": {
    "available": true
  },
  "children": [
    {"age": 5},
    {"name": "Bob", "age": null},
    {"name": "Judit"}
  ]
}'

