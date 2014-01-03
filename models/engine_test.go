package models

import "testing"

func init() {
	CreateDb()
}

var (
	user = "work"
	pass = "krow"
)

func TestSyncRecord(t *testing.T) {
	r := new(Record)
	r.Hostname = "example.com"
	r.User = user
	r.Pass = pass

	if err := r.Sync(); err != nil {
		t.Fatal(err)
	}

	r.Pass = ""
	var rc *Record
	var err error
	if rc, err = GetRecord(r.Hostname, user); err != nil {
		t.Fatal(err)
	}

	if rc.Pass != "krow" {
		t.Error("expect krow, but not this")
	}
	t.Log(rc)
}
