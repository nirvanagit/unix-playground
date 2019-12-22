## Sharing sockets across users

Given we have two users, ironman from group admin, and foo from group play,
we want ironman to create a socket and enable foo to connect to it.

Requirements to make that possible
- sock file should be created in a directory which both users can share and have write permissions to. This can be done by using the sticky bit, which enables users to create file, but delete only those files which they created

- set correct permissions for sock file
  + provide write access to all. To enable this we need two things, disable umask, and give rw permissions to all users. Ref - [stackoverflow](https://stackoverflow.com/questions/35424970/unix-socket-permissions-linux)

Docker run

```
docker run -it centos:latest sh
```

Set umask to 0

```
umask 0
```
Create two users

```
sh-4.4# groupadd admin
sh-4.4# useradd ironman -G admin

sh-4.4# groupadd play
sh-4.4# useradd foo -G play
```

Create a directory `/socket` directory and set its sticky bit

```
sh-4.4# mkdir /socket 
sh-4.4# chmod o+t socket

# grant execute permissions to all
sh-4.4# chmod a+rwx socket

sh-4.4# ls -ld socket
drwxrwxrwt 2 root root 4096 Dec 22 19:58 socket
```

Login as ironman and create socket

```
sh-4.4# su ironman
[ironman@1f0e095ac699 /]$ cd socket
[ironman@1f0e095ac699 socket]$ socat UNIX-LISTEN:bar.sock -
```

In another shell grant rw access to all on /socket/bar.sock

```
[ironman@1f0e095ac699 socket]$ chmod a+rw bar.sock
```

Login as user foo and connect to the socket

```
[foo@1f0e095ac699 socket]$ ncat -U bar.sock -v
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Connected to bar.sock.
hello world
```
