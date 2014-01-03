## gossh 

Beta released (2013-1-2)
*Developing*

[![Build Status](https://drone.io/github.com/shxsun/gossh/status.png)](https://drone.io/github.com/shxsun/gossh/latest)

gossh is a command which wrapper for ssh and [sshpass](https://linuxtoy.org/archives/sshpass.html).

Program use basic C/S structure. passwords can be stored in local db.

### gossh can de what?
1. quickly connect to a host, password will store in local database.
2. *todo: support alias of hostname(very usefull if your hostname is very long)*

### How to use
1. Download `sshpass` from <http://sourceforge.net/projects/sshpass/>, and add it to PATH
2. start gossh server: `nohup ./gossh -server &>/gossh.log &`
3. use gossh just like ssh. for example: `gossh root@example.com`

password required only for the first time.

### The working flow.
1. client login to server, (todo: server check auth)
2. client send user(eg: root), host(eg: example.com) to server
3. server search hostname which user=root and hostname matches example.com
4. server send back password
5. client call sshpass and ssh to connect host

### How to join develop
There are still lot of things to do. Even throuth the first version is released. [todo](doc/todo.md)

download thrift if you want to change protocal: <http://mirror.esocc.com/apache/thrift/0.9.1/thrift-0.9.1.tar.gz>

use `thrift --gen go rpc.thrift` to generate go code.

*Document*

* [server api document](doc/api.md)
* [server storage](doc/storage.md)
