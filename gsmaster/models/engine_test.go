package models

import "testing"

func init() {
	CreateDb()
}

func TestSyncRecord(t *testing.T) {
	r := new(Record)
	r.Hostname = "example.com"
	r.User = "work"
	r.Pass = "krow"

	err := r.Sync()
	if err != nil {
		t.Fatal(err)
	}

	r.Pass = ""
	err = r.Get()
	if err != nil {
		t.Fatal(err)
	}
	if r.Pass != "krow" {
		t.Error("expect krow, but not this")
	}
	t.Log(r)
}
