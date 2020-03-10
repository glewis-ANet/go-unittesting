# go-unittesting

Example unit testing project for GoLang

This project demonstrates some unit testing features in GoLang using

* Testify (assertion and mocking package)
* Mockery (mock generator)

Relevant commands:

```
mockery -dir interfaces -name Serialiser
```

Command used to generate mocks (in this case the mock for the Serialiser interface)

```
go test -v -coverprofile=coverage.out ./...
```

Run the tests and generate code coverage

```
go tool cover -html=coverage.out
```

View the coverage report in HTML format
