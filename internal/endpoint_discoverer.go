package internal

import (
	"net"

	"github.com/nickburkhart/go-gesclient/tasks"
)

type EndpointDiscoverer interface {
	DiscoverAsync(ipEndpoint net.Addr) *tasks.Task
}

type staticEndpointDiscoverer struct {
	task *tasks.Task
}

func NewStaticEndpointDiscoverer(ipEndpoint net.Addr, isSsl bool) *staticEndpointDiscoverer {
	if ipEndpoint == nil {
		panic("ipEndpoint is nil")
	}
	var nodeEndpoints *NodeEndpoints
	if isSsl {
		nodeEndpoints = NewNodeEndpoints(nil, ipEndpoint)
	} else {
		nodeEndpoints = NewNodeEndpoints(ipEndpoint, nil)
	}
	task := tasks.New(func() (interface{}, error) {
		return nodeEndpoints, nil
	})
	return &staticEndpointDiscoverer{
		task: task,
	}
}

func (d *staticEndpointDiscoverer) DiscoverAsync(ipEndpoint net.Addr) *tasks.Task {
	return d.task
}
