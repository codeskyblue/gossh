## Storage of server

For a rds

	tblUser
		id(int64) username(string) password(string) view_count(int64) born_time(time) type(enum:user, group)
	tblGroup
		user_id(int64) group_user_id(int64)
	tblHost
		id(int64) hostname(string) view_count(int64)
	tblToken
		user_id(int64) token(string) expire(time)

		
