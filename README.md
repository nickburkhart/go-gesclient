# Event Store Go Client

A [Go](https://golang.org/) port of the .Net client for [Event Store](https://eventstore.org/).

## Status

This project is considered at an alpha stage and should'nt be used in production.

## Warning

API is still under development and could change.

## Requirements

- Go 1.4+

## License

MIT. See [LICENSE](https://github.com/jdextraze/go-gesclient/blob/master/LICENSE).

## Implemented

* Writing to a stream
* Reading a single event
* Reading a stream forwards
* Reading a stream backwards
* Reading all events forwards
* Reading all events backwards
* Volatile subscriptions
* Persistent subscription
* Deleting stream
* Cluster connection
* Global authentication
* Get/Set stream metadata
* Set system settings

## TODO

* Complete unit and integration tests
* Benchmarks
* Transaction

## External tools

* [Robo](https://github.com/tj/robo) (`go get github.com/tj/robo`)
* [Docker](https://www.docker.com/get-docker)

## Examples

For examples, look into `examples`. You will need an instance of event store to be running to try it.
I suggest using [Docker](https://docker.com/) with [Event Store Docker Container](https://hub.docker.com/r/eventstore/eventstore/).

## Other languages client

* [.Net](https://github.com/EventStore/EventStore) (Official)
* [Java / Scala](https://github.com/EventStore/EventStore.JVM) (Official)
* [Node.JS](https://github.com/nicdex/node-eventstore-client)
