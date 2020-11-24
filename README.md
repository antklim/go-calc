# go-calc
Basic calculator service in Go

## Usage
To build service run `make build`.  
`calc` binary will be created in `./build/calc`. To launch service run `./build/calc`.

To build and start service in a container, run the following commands:
```
$ make services-build # build service docker image
$ make services-up    # launch service container
$ make services-down  # stop service container
```

Under the hood all `service-*` commands use `docker-compose` commands, so could be replaced with:
```
$ docker-compose build
$ docker-compose up
$ docker-compose down
```

## Testing
To execute unit-tests run `make test`

To execute integration tests run:
```
$ make services-build
$ make services-up
$ cd test
$ npm t
or
$ newman run postman_collection.json
```