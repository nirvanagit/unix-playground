### Process group id
ID of the parent to which the process belongs

### Background
When you open a terminal session, you create a new process, and anything you run from that terminal will be it's child process.

### How to traverse process id to the pid 1 in a ssh session?

Login to a machine using ssh, and run the following
```
# get the process id of the login shell
[foo@bar processes]$ echo $$
8482

# run sleep in background
[foo@bar processes]$ sleep 120 &
[1] 23957

# capture ppid, pgid
# and we notice that PPID of the process running sleep is the pid of our login shell
[foo@bar processes]$ ps -p 23957 -o pid,ppid,pgid,command
  PID  PPID  PGID COMMAND
23957  8482 23957 sleep 120

# login shell command is `bash`
[foo@bar processes]$ ps -p 8482 -o pid,ppid,pgid,command
  PID  PPID  PGID COMMAND
 8482  8481  8482 -bash

# parent of our login shell pid is the sshd server process
[foo@bar processes]$ ps -p 8481 -o pid,ppid,pgid,command
  PID  PPID  PGID COMMAND
 8481  8477  8477 sshd: foo@pts/1

# parent of sshd shows that it is running as a privilege separation thread
[foo@bar processes]$ ps -p 8477 -o pid,ppid,pgid,command
  PID  PPID  PGID COMMAND
 8477  7998  8477 sshd: foo [priv]

# finally the parent of our sshd thread is pid 1
[foo@bar processes]$ ps -p 7998 -o pid,ppid,pgid,command
  PID  PPID  PGID COMMAND
 7998     1  7998 /usr/sbin/sshd -D

# as pid 1 does not have a parents, its ppid is set to _0_
[foo@bar processes]$ ps -p 1 -o pid,ppid,pgid,command
  PID  PPID  PGID COMMAND
    1     0     1 /usr/lib/systemd/systemd --system --deserialize 24
```
