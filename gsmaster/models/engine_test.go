package models

import "testing"

func init() {
	CreateDb()
}

func TestAddRecord(t *testing.T) {
	r := new(Record)
	r.Hostname = "example.com"
	r.User = "work"
	r.Pass = "krow"

	err := r.Sync()
	if err != nil {
		t.Fatal(err)
	}

	r.Pass = "wwwwww"
	err = r.Get()
	if err != nil {
		t.Fatal(err)
	}
	if r.Pass != "krow" {
		t.Error("expect krow, but not this")
	}
	t.Log(r)
	//err := SyncRecord(r)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//r, _ = GetRecord(r.Hostname, r.User)
	//t.Log(r)
}

/*
func TestAddNew(t *testing.T) {
	hu := new(HostUser)
	hu.Host.Hostname = "nihao"
	hu.User.Username = "ssx"
	hu.User.Password = "mypass"
	_, err := Engine.InsertOne(hu)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCassade(t *testing.T) {
	return
	//host := new(models.Host)
	//	host.Hostname = "xyzdas"
	//	host.Alias = "abc"
	//	n, err := models.Engine.InsertOne(host)
	hostname := "291abcfe"
	host := new(Host)
	host.Hostname = hostname
	host.Alias = "a"
	affed, err := Engine.InsertOne(host)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(affed)

	user := new(User)
	user.Username = "work"
	user.Password = "krow"
	affed, err = Engine.InsertOne(user)
	if err != nil {
		t.Error(err)
	}
	log.Println(affed)

	hu := new(HostUser)
	hu.Host = *host
	hu.User = *user
	affed, err = Engine.InsertOne(hu)
	if err != nil {
		t.Error(err)
	}
	log.Println(affed)

	host = new(Host)
	Engine.Where("hostname=?", hostname).Get(host)
	fmt.Println(host)
	hu = new(HostUser)
	hu.Host.Hostname = hostname
	Engine.Get(hu)
	//Engine.Where("host_id=?", host.Id).Get(hu)
	fmt.Println(hu)
}*/
