gossh *Developing*
=====

This is not ssh. This is a help tool for ssh.

use gossh can quickly connect to a system, and scp file to any where you want.

gossh has two parts, cloud-server and client(name gossh)

### How to join
all develops in todo [TODO](todo.md)

### How to use
1. First you need to store `hostname, username, password` to cloud[?](cloud.md)

for example: cloud stores example-vps.com, root, 123456

	# client use such command to connect to example-vps.com
	gossh root@example echo "hello world"

Some things will happen

1. client login cloud server, check auth
2. client send user(root), hostAlias(example) to cloud server
3. server search hostname which user=root, hostname matches example.
4. server send back alalible list of (hostname, password)
5. client call ssh, scp or rsync to do commands (use sshpass of simply gosshpass)

### Other ideas
for a group use, give them a group-account.
for personal use, give them a single-account and group is group-account

so personal can use itself conf and also the group conf.

