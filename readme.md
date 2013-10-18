# Gotham

HTTP API for a comic book app.

## API Documentation

### Home resource

The home resource is the only resource that an API client should have hardcoded. It represents the entrypoint to the Gotham Hypermedia API.

Home is represented using the json-home Media Type, described in http://tools.ietf.org/html/draft-nottingham-json-home.

*Request*
```HTTP
GET / HTTP/1.1
Host: gotham
Accept: application/json-home
```

*Response*
```HTTP
HTTP/1.1 200 OK
Content-Type: application/json-home
```
```JSON
{
	"resources": {
		"http://gotham/rel/featured": { "href": "/featured" }
	}
}
```

### Featured resource

*Request*
```HTTP
GET /featured HTTP/1.1
Host: gotham
Accept: application/comics+json
```

*Response*
```HTTP
HTTP/1.1 200 OK
Content-Type: application/comics+json
```
```JSON
{
	"set": [
		{
			"type": "series",
			"title": "Prophet",
			"links": {
				"via": {"href": "/series/1"}
			},
		},
		{
			"type": "series",
			"title": "Jupiter's Legacy",
			"links": {
				"via": {"href": "/series/2"}
			},
		}
	]
}
```