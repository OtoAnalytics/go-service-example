# go-service-example
Example web service in Go using the Martini library, with some test endpoints and some assets.

Run locally with "go run server.go" and it will listen on port 2003.

Test `localhost:2003/` to see "Hello world".

Test `localhost:2003/locations/:id` to see "Hello location id" where "id" is whatever parameter you provide.

Test `localhost:2003/hello.html` to see the "public/hello.html" page as an example of a static asset.

Test `localhost:2003/asset.html` to see the "assets/asset.html" page as an example of mapping another asset directory.

### Martini

See documentation on the Martini library here: https://github.com/go-martini/martini


