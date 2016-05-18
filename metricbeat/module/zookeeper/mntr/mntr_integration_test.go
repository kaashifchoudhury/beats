// +build integration

package mntr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/common"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
	"github.com/elastic/beats/metricbeat/module/zookeeper"
)

func TestFetch(t *testing.T) {
	f := mbtest.NewEventFetcher(t, getConfig())
	event, err := f.Fetch()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)

	// Check values
	version := event["version"].(string)
	avgLatency := event["latency"].(common.MapStr)["avg"].(int)
	maxLatency := event["latency"].(common.MapStr)["max"].(int)
	numAliveConnections := event["num_alive_connections"].(int)

	assert.Equal(t, version, "3.4.8--1, built on 02/06/2016 03:18 GMT")
	assert.True(t, avgLatency >= 0)
	assert.True(t, maxLatency >= 0)
	assert.True(t, numAliveConnections > 0)

	// Check fields
	assert.Equal(t, 15, len(event))
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "zookeeper",
		"metricsets": []string{"mntr"},
		"hosts":      []string{zookeeper.GetZookeeperEnvHost() + ":" + zookeeper.GetZookeeperEnvPort()},
	}
}
