// client_test.go
package main

import (
	"testing"
)

func TestSplitUserHost(t *testing.T) {
	user, host, err := SplitUserHost("work@example.com")
	t.Log(user, host)
	if err != nil {
		t.Error(err)
	}
	if user != "work" || host != "example.com" {
		t.Errorf("expect user:work host:example.com, but got user:%s host:%s", user, host)
	}
}

func TestSplitUserHost2(t *testing.T) {
	user, host, err := SplitUserHost("example.com")
	t.Log(user, host)
	if err != nil {
		t.Error(err)
	}
	if user == "" || host != "example.com" {
		t.Errorf("expect user:<curr> host:example.com, but got user:%s host:%s", user, host)
	}
}
