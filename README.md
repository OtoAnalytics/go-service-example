# go-service-example
Example web service in Go using the Martini library, with some test endpoints and some assets.

### Running

Run locally with "go run server.go" and it will listen on port 3000. NOTE if you have an environment variable PORT then it will listen on that port instead by default.

	Browse to "/" to see "Hello world".

	Browse to "/locations/:id" to see "Hello location id" where "id" is whatever parameter you provide.

	Browse to "/hello.html" to see the "public/hello.html" page as an example of a static asset.

	Browse to "/asset.html" to see the "assets/asset.html" page as an example of mapping another asset directory.

### Building a (tiny!) Docker container

In the working directory after cloning this repository:

1. Make sure you have docker-machine up and running, and know what your `docker-machine ip` address is.
1. `go get github.com/go-martini/martini`
1. `godep save -r`
1. `docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:1.5 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o go-service-example'`
1. `docker build -t go-service-example .`
1. `docker run --rm -it -p 3000:3000 go-service-example`
1. Browse to your docker-machine IP address port 3000 and you should see "Hello world"!

If you do a `docker images` and look for this one, it will only be about 5.5MB.

### Martini

See documentation on the Martini library here: https://github.com/go-martini/martini

