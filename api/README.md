#  ForgeResponse
Security Incident Management Platform API

# Development

To get started with local development you will need a few prerequisites:
* [Golang >= 1.19.x](https://go.dev/dl/)
* [Docker](https://www.docker.com/)
* [Venom](https://github.com/ovh/venom)

Create a .env file and populate it with the following:

```bash
audience="dev app audience"
issuerUrl="https://dev-auth0-tenant/"
```

Once the dependencies have been installed and the .env file created run the following command to start the API server and a MongoDB container:

```bash
docker compose --env-file .env up
```

## Tests

The code base uses both unit tests and E2E tests.

Unit tests can be run from the root level of the source code by running `go test ./...`

E2E tests are built using Venom which allows the tests to be declared in yaml files. To run the E2E tests perform the following steps:

* Run `docker compose up`
* In another terminal window navigate to tests/e2e
* Run `venom run` this will run all the tests. If you want to run specific tests then use `venom run filename.yml`

