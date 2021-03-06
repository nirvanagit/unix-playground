## Difference between real, effective and saved user ids

### Real user id

This is the user id who created the process.
```
# check UID of the logged in user
[foo@bar ~]$ grep $LOGNAME /etc/passwd
foo:x:85478:100:foo:/home/foo:/bin/bash
[foo@bar ~]$

# execute a command
[foo@bar ~]$ sleep 120 & ps aux | grep 'sleep'
[1] 18129
foo   18127  0.0  0.0 217044   816 ?        S    22:54   0:00 sleep 1
foo   18129  0.0  0.0 217044   716 pts/1    S    22:54   0:00 sleep 120
foo   18131  0.0  0.1 221900  1056 pts/1    R+   22:54   0:00 grep --color=auto sleep

# get the UID and GID of the above process
[foo@bar ~]$ stat -c "%u %g" /proc/18129
85478 100
[foo@bar ~]$
```

### Effective User id

If there exists a binary which requires root access to run, but should be able to run by a normal user,
then this binary will have its setuid bit set

What this means is, that a non-privileged user will be able to execute a privileged command

For instance, lets look at the mount command
```
[foo@bar ~]$ ls -la /usr/bin/mount
-rwsr-xr-x. 1 root root 66432 Sep 21 09:45 /usr/bin/mount
[foo@bar ~]$
```

It has its set-uid bit set (notice the __s__ instead of __x__)
The owner and group of /usr/bin/mount is root. This is because mount command runs some privileged process and the kernel demands root privilege for that.

This is where EUID and EGID comes into play. When an underprivileged user executes the mount command, the process changes its effective user id (EUID) from the default
RUID to the owner of this special binary executable file, which in this case is __root__.

The kernel makes the decision whether this process has the privilege by looking at the EUID of the process.

Setting SGID, is same as setting SUID for files, but differs for directories.

In the following exercise, we have two users, foo and bar who belong to the group 'play'. foo and bar want to collaborate on a folder called /project.

Once the admin configures this folder and assigns the groupid of the folder to play, when foo creates a file /project/foo-file the ownership is foo(user) and foo(group), similar behavior when
bar creates a file.

What we want to do is, whenever either of them creates a file, the group ownership is set to 'play'.

```
[root@cd5b15a4228e /]# whoami
root
[root@cd5b15a4228e /]# mkdir project
[root@cd5b15a4228e /]# ls -ld project
drwxr-xr-x 2 root root 4096 Dec 20 02:54 project
[root@cd5b15a4228e /]# chgrp -R play project
[root@cd5b15a4228e /]# ls -ld project
drwxr-xr-x 2 root play 4096 Dec 20 02:54 project
[root@cd5b15a4228e /]# chmod g+rwx project
[root@cd5b15a4228e /]# ls -ld project
drwxrwxr-x 2 root play 4096 Dec 20 02:54 project
[root@cd5b15a4228e /]# exit
exit
sh-4.4# su foo
[foo@cd5b15a4228e /]$ touch project/foo-file
[foo@cd5b15a4228e /]$ exit
exit
sh-4.4# su bar
[bar@cd5b15a4228e /]$ touch project/bar-file
[bar@cd5b15a4228e /]$ ls -la project
total 8
drwxrwxr-x 2 root play 4096 Dec 20 02:55 .
drwxr-xr-x 1 root root 4096 Dec 20 02:54 ..
-rw-rw-r-- 1 bar  bar     0 Dec 20 02:55 bar-file
-rw-rw-r-- 1 foo  foo     0 Dec 20 02:55 foo-file
[bar@cd5b15a4228e /]$
[bar@cd5b15a4228e /]$ exit
exit
sh-4.4# su root
[root@cd5b15a4228e /]#
[root@cd5b15a4228e /]#
[root@cd5b15a4228e /]# chmod g+s project
[root@cd5b15a4228e /]# exit
exit
sh-4.4# su foo
[foo@cd5b15a4228e /]$ touch project/foo-file-2
[foo@cd5b15a4228e /]$ ls -la project
total 8
drwxrwsr-x 2 root play 4096 Dec 20 02:56 .
drwxr-xr-x 1 root root 4096 Dec 20 02:54 ..
-rw-rw-r-- 1 bar  bar     0 Dec 20 02:55 bar-file
-rw-rw-r-- 1 foo  foo     0 Dec 20 02:55 foo-file
-rw-rw-r-- 1 foo  play    0 Dec 20 02:56 foo-file-2
[foo@cd5b15a4228e /]$
```

### Sticky bit

This bit should have always been called the "restricted deletion bit" given that's what it really connotes.
When this mode bit is enabled, it makes a directory such that users can only delete files & directories within it that they are the owners of.

- Sticky bit is used on shared directories.
- It is useful for shared directories such as `/var/tmp/` and `/tmp` because users can create files, read and execute files
owned by other users, but are not allowed to remove files owned by other users.
- For example if user bob creates a file named /tmp/bob, other user tom can not delete this file even when the /tmp directory has permission of 777. If sticky bit is not set then tom can delete /tmp/bob, as the /tmp/bob file inherits the parent directory permissions.
– root user (Off course!) and owner of the files can remove their own files.

If `t` tag is set in a directory, it means its sticky bit is set.

```
[foo@bar ~]$ ls -ld /tmp
drwxrwxrwt. 4 root root 107 Dec 19 23:15 /tmp
[foo@bar ~]$
```

To set a sticky bit use `+t` and `-t` to remove.
```
$ chmod o-t dir1
$ ls -l
total 8
drwxr-xr-x 2 root root 4096 Aug 19 03:08 dir1
$ chmod o+t dir1
$ ls -l
total 8
drwxr-xr-t 2 root root 4096 Aug 19 03:08 dir1
```

```
- T refers to when the execute permissions are off.
- t refers to when the execute permissions are on.
```

### References

- [stackoverflow](https://stackoverflow.com/questions/32455684/unix-linux-difference-between-real-user-id-effective-user-id-and-saved-user)
- [stackoverflow](https://unix.stackexchange.com/questions/79395/how-does-the-sticky-bit-work)
- [linoxide](https://linoxide.com/how-tos/stickbit-suid-guid/)
- [thegeekdiary.com](https://www.thegeekdiary.com/what-is-suid-sgid-and-sticky-bit/)
