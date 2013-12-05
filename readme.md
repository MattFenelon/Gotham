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
			"links": [
				{"rel":"self", "href": "http://gotham/series/random"},
				{"rel":"seriesimage", "href": "http://gotham/pages/random/0.jpg"},
				{"rel":"promotedbook", "href": "http://gotham/books/random"}
			],
		},
		{
			"title": "Jupiter's Legacy",
			"links": [
				{"rel":"self", "href": "http://gotham/series/random"},
				{"rel":"seriesimage", "href": "http://gotham/pages/random/0.jpg"},
				{"rel":"promotedbook", "href": "http://gotham/books/random"}
			],
		}
	]
}
```

### Series resource

*Request*
```HTTP
GET /series/{opaque} HTTP/1.1
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
	"title": "Saga",
	"books": [
		{
			"title": "Saga 13",
			"publishedDate": "2013-08-14T00:00:00Z", // RFC 3339 format
			"writtenBy": "Brian K. Vaughan",
			"artBy": "Fiona Staples",
			"blurb": "THE SMASH-HIT, CRITICALLY ACCLAIMED SERIES RETURNS!\nNow that you've read the first two bestselling collections of SAGA, you're all caught up and ready to jump on the ongoing train with Chapter Thirteen, beginning an all-new monthly sci-fi/fantasy adventure, as Hazel and her parents head to the planet Quietus in search of cult romance novelist D. Oswald Heist.",
			"links": [
				{"rel":"self", "href": "http://gotham/books/{opaque}"},
				{"rel":"bookimage", "href": "http://gotham/pages/{opaque}"}
			]
		},
		{
			"title": "Saga 12",
			"publishedDate": "2013-04-10T00:00:00Z",
			"writtenBy": "Brian K. Vaughan",
			"artBy": "Fiona Staples",
			"blurb": "Prince Robot IV makes his move.",
			"links": [
				{"rel":"self", "href": "http://gotham/books/{opaque}"},
				{"rel":"bookimage", "href": "http://gotham/pages/{opaque}"}
			]
		}
	]
}
```

### Comic resource

#### Add a comic

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

#### Get a comic

*Request*
```HTTP
GET /books/{random} HTTP/1.1
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
	"links":[
		{"rel":"item","href":"http://gotham/pages/random/0.jpg"},
		{"rel":"item","href":"http://gotham/pages/random/1.jpg"},
		{"rel":"item","href":"http://gotham/pages/random/2.jpg"},
		{"rel":"item","href":"http://gotham/pages/random/3.jpg"}
	]
}
```