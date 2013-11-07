# Gotham

HTTP API for a comic book app.

## API Documentation

### Home resource

*Request*
```HTTP
GET / HTTP/1.1
Host: gotham
Accept: application/json
```

*Response*
```HTTP
HTTP/1.1 200 OK
Content-Type: application/json
```
```JSON
{
	"series": [
		{
			"title": "Prophet",
			"links": {
				"seriesimage": {"href": "http://gotham/pages/random/0.jpg"}
			},
		},
		{
			"title": "Jupiter's Legacy",
			"links": {
				"seriesimage": {"href": "http://gotham/pages/random/0.jpg"}
			},
		}
	]
}
```

### Comic resource

The comic resource can be used to add comics.

Requests are formed as multipart/form-data. The first part is expected to be the metadata for the comic represented as application/json with "metadata" as the field-name. The subsequent parts should be the page images for the comic, which must have a field-name of "page". The page images must be in the order by which they are to be displayed. The images are stored under the filename specified in the filename parameter of the part. The only page image format currently supported is image/jpeg.

A comic must have at least 1 page image.

*Request*
```HTTP
POST /books HTTP/1.1
Host: gotham
Accept: application/json
Content-Type: multipart/form-data; boundary=Any-ASCII-string
```
```

--Any-ASCII-string
Content-Disposition: form-data; name="metadata"
Content-Type: application/json

{
	"seriesTitle": "Prophet",
	"title": "Prophet 31",
}
--Any-ASCII-string
Content-Disposition: form-data; name="page"
Content-Type: image/jpeg

<Binary content goes here>
--Any-ASCII-string
Content-Disposition: form-data; name="page"
Content-Type: image/jpeg

<Binary content goes here>
--Any-ASCII-string
Content-Disposition: form-data; name="page"; filename="0.jpg"
Content-Type: image/jpeg

<Binary content goes here>
--Any-ASCII-string--
```

*Response*
```HTTP
HTTP/1.1 204 No Content
```
