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
		"{docHost}/rel/featured": { "href": "/featured" }
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
				"via": {"href": "/series/1"},
				"{docHost}/rel/seriesimage": {"href": "/images/1.jpg"}
			},
		},
		{
			"type": "series",
			"title": "Jupiter's Legacy",
			"links": {
				"via": {"href": "/series/2"},
				"{docHost}/rel/seriesimage": {"href": "/images/2.jpg"}
			},
		}
	]
}
```

### Comic resource

The comic resource can be used to add comics.

Requests are formed as multipart/mixed. The first part is expected to be the metadata for the comic represented as application/json. The subsequent parts are the page images for the comic. The page images must be in the order by which they are to be displayed. The only page image format currently supported is image/jpeg.

A comic must have at least 1 page image.

*Request*
```HTTP
POST /books HTTP/1.1
Host: gotham
Accept: application/json
Content-Type: multipart/mixed; boundary=Any-ASCII-string
```
```HTTP

--Any-ASCII-string
Content-Type: application/json

{
	"seriesTitle": "Prophet",
	"title": "Prophet 31",
}
--Any-ASCII-string
Content-Disposition: attachment; filename="00.jpg"
Content-Type: image/jpeg

<Binary content goes here>
--Any-ASCII-string
Content-Disposition: attachment; filename="01.jpg"
Content-Type: image/jpeg

<Binary content goes here>
--Any-ASCII-string
Content-Disposition: attachment; filename="nn.jpg"
Content-Type: image/jpeg

<Binary content goes here>
--Any-ASCII-string--
```

*Response*
```HTTP
HTTP/1.1 204 No Content
```