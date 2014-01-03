## Storage of server

For a rds

	simple structure
	-----------------------------------
	tblHost
		hostname: string
		user: string
		pass: string
		view_count: int64
		
	
	safe structure
	-----------------------------------
	tblUser
		id: int64 
		username: string
		hash_login_password: string
		view_count: int64
		born_time: time
		type: enum{user, group}
		aes_info_password: string
	tblGroup
		user_id: int64
		group_user_id: int64
	tblHostUser
		id: int64
		user_id: int64
		username: string
		aes_password: string
	tblHost
		hostname: string
		host_user_id: int64
		view_count: int64
	tblVirtPassword
		id: int64
		user_id: int64
		aes_password: string

## Links
* <http://www.zendstudio.net/archives/php-google-authenticator/>
		
