# JSON Service

This repo contains a very simple service which processes an uploaded json file and stores the extracted information in a database.
For simplicity, we use in-memory storage.

## Requirements

- [Go](https://golang.org/doc/install) >= Go 1.19
- [GNU Make](https://www.gnu.org/software/make/)
- [Docker](https://docs.docker.com/engine/install)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Endpoints

The service has one single endpoint but accepts two HTTP methods

```
GET /ports
POST /ports
```

The `POST` endpoint accepts as a paramter an uploaded json file, something like the one bellow:

```json
{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }
}
...
```

## Getting started

To run this service in the background just call

```shell
make up
```

## Testing

To run all tests just call

```shell
make tests
```

## Development

* **setup** - setup development tools (this only downloads the linter binary)
* **clean** - delete all binaries downloaded during the setup stage
* **lint** - run configured linters for static code analysis
* **tests** run all tests
* **up** start all docker containers
* **down** stop all docker containers
