## gossh 
*Developing*

This is not ssh. This is a help tool for ssh.

### gossh can de what?
1. quickly connect to a host(of course you donot have to type password), and then do shell commands or copy sths.
2. TODO: support alias of hostname(very usefull if your hostname is very long)

gossh has two parts, server and client.

### How to use gossh
##### 1. Start gossh server
	./gossh -server

##### 2. Use gossh instead of ssh
The first time you will need to type the password, but the second time will not.

	./gossh work@example.com

### The working flow.
1. client login to server, (TODO: server check auth)
2. client send user(eg: root), host(eg: example.com) to server
3. server search hostname which user=root and hostname matches example.com
4. server send back password
5. client call sshpass to connect host

### How to join develop
All contributes can be found in [todo](doc/todo.md)

download thrift if you want to change protocal: <http://mirror.esocc.com/apache/thrift/0.9.1/thrift-0.9.1.tar.gz>

use `thrift --gen go rpc.thrift` to generate go code.

*Document*

* [server api document](doc/api.md)
* [server storage](doc/storage.md)
