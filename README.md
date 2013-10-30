gossh *Developing*
=====

This is not ssh. This is a help tool for ssh.

gossh has two parts, cloud-server and client(name gossh)

### How to join
all develops in todo [TODO](todo.md)

## Cloud Server
### Work flow
1. store `hostname, username, password` to cloud
2. manage client query

client query usally `accountName, accountPass, shortHostname, user`

for example:

	# Server stores example-vps.com, root, 123456
	# Client goes that
	gossh root@example echo "hello world"

Some things will happen

1. client login cloud server, check auth
2. client send user(root), hostAlias(example) to cloud server
3. server search hostname which user=root, hostname matches example.
4. server send back alalible list of (hostname, password)
5. client call ssh, scp or rsync to do commands (use sshpass of simply gosshpass)

### How to manage infos to cloud server
1. use a browser to login with https://
2. do add, update, delete

### Other ideas
for a group use, give them a group-account.
for personal use, give them a single-account and group is group-account

so personal can use itself conf and also the group conf.
