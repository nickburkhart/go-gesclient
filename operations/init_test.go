package operations_test

import (
	"net/url"
	"time"

	"github.com/nickburkhart/go-gesclient"
	"github.com/nickburkhart/go-gesclient/client"
)

var es client.Connection

func init() {
	ensureConnection()
}

func ensureConnection() {
	if es != nil {
		return
	}

	var err error

	uri, _ := url.Parse("tcp://127.0.0.1:1113/")
	es, err = gesclient.Create(client.DefaultConnectionSettings, uri, "benchmark")
	if err != nil {
		panic(err)
	}
	es.Disconnected().Add(func(event client.Event) error { panic("disconnected") })
	es.ConnectAsync().Wait()
	time.Sleep(100 * time.Millisecond)
}
