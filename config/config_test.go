package config

import "testing"

func TestDummy(t *testing.T) {
}

func TestServerAddress(t *testing.T) {
	var configuration Configuration
	configuration.Host = "server.com"
	configuration.Port = 1234
	if configuration.ServerAddress() != "server.com:1234" {
		t.Error("ServerAddress expected server.com:1234 but got", configuration.ServerAddress())
	}
}

func TestServerURL(t *testing.T) {
	var configuration Configuration
	configuration.Host = "server.com"
	configuration.Port = 1234
	if configuration.ServerURL() != "http://server.com:1234" {
		t.Error("ServerAddress expected http://server.com:1234 but got", configuration.ServerURL())
	}
}
