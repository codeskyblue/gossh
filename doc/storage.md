## Storage of server

For a rds

	tblUser
		id(int64) username(string) password(string) view_count(int64) born_time(time) type(enum:user, group)
	tblGroup
		user_id(int64) group_user_id(int64)
	tblHostUser
		id(int64) user_id(int64) username(string) password(string)
	tblHost
		hostname(string) host_user_id(int64) view_count(int64)
	tblToken
		user_id(int64) token(string) expire(time)

		
