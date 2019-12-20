### List all groups
```
[foo@bar ~]$ getent group
tcpdump:x:1:
...
```

### List users in a group
```
[foo@bar ~]$ sudo lid -g users
 games(uid=12)
 scanserf(uid=33)
 foo(uid=22)
[foo@bar ~]$
```

### List group membership of a user
```
[foo@bar ~]$ groups foo
foo : users
[foo@bar ~]$
```
