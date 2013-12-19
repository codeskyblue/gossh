namespace go rpc

struct Data {
	1: required string hostname;
	2: required string password;
	3: required string username;
}

service Gs {
	Data LookHost(1:string hostname, 2:string username);
	bool SyncHost(1:Data data);
}