## Zombies and orphaned processes

### Zombies

Give a process A, whenever it creates a child process B, it is A's reponsibility to wait for B to exit, take its
exit details, such as exit code etc. To do this A performs a __wait__ operation.

In psuedocode

```
1. create A
2. A creates child B
3. A waits for child B
4. B exits
5. A gets B's exit details
6. A exits
```
In a scenario where B exits before A gets a chance to wait for it, there is no way A can consistently get B's information.
To solve this problem, the kernel turns every child process into a zombie. Zombies are what they are in real life (they cannot be killed, :P, even with a SIGKILL).

This allows A to eventually to a wait for B, and get information such as child's process ID, termination status, and resource usage statistics.

Given that every child will always be turned into a zombie until the parent performs a wait operation, what happens when parent fails to perform wait? In this case the child will indefinitely remain
in the kernel's process table.

E.g.
Create a file called `create_zombie.sh`
```
#!/bin/sh

sleep 1 & exec sleep 30
```

Before invoking script run

ps aux
```
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.0   4452  1584 pts/0    Ss   00:26   0:00 sh
root        50  0.0  0.0   4452   660 pts/1    Ss   00:41   0:00 sh
root        63  0.0  0.0  15576  2204 pts/1    R+   00:59   0:00 ps aux
#
```

top # to capture there are no zombies
```
top - 01:00:29 up  2:42,  0 users,  load average: 0.00, 0.00, 0.00
Tasks:   3 total,   1 running,   2 sleeping,   0 stopped,   0 zombie
...
```

execute it
```
# sh create_zombie.sh
```

Check
ps aux
```
# ps aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.0   4452  1584 pts/0    Ss   00:26   0:00 sh
root        50  0.0  0.0   4452  1504 pts/1    Ss   00:41   0:00 sh
root        71  0.0  0.0   4352   652 pts/0    S+   01:01   0:00 sleep 10
root        72  0.0  0.0      0     0 pts/0    Z+   01:01   0:00 [sleep] <defunct>
root        75  0.0  0.0  15576  2108 pts/1    R+   01:02   0:00 ps aux
```

Creation of zombie
```
# top
top - 01:01:28 up  2:43,  0 users,  load average: 0.00, 0.00, 0.00
Tasks:   3 total,   1 running,   2 sleeping,   0 stopped,   1 zombie
```

What happened here?
1. From within the shell script we ran `sleep 1 &` in background.
2. Now normally this process would be inherited by the process created when we invoked `sh create_zombie.sh`, but what instead happened here
was, we ran `exec sleep 10`. Running sleep with __exec__ took over the PID of `sh create_zombie.sh`, and it won't know when sleep exited.

Just to show what exec did, I've remove __exec__ from `create_zombie.sh`.
```
# sh create_zombie.sh
#

# # in another shell
# ps aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.0   4452  1584 pts/0    Ss   00:26   0:00 sh
root        50  0.0  0.0   4452  1504 pts/1    Ss   00:41   0:00 sh
root        79  0.0  0.0   4452   740 pts/0    S+   01:17   0:00 sh create_zombie.sh
root        81  0.0  0.0   4352   644 pts/0    S+   01:17   0:00 sleep 10
root        82  0.0  0.0  15576  2144 pts/1    R+   01:17   0:00 ps aux
#
```

This time we see that `sh create_zombie.sh` and `sleep 10` have different PIDs.

Note: running sleep with __exec__ is not making `sleep 1 &` an orphan, because it was originally called from within the script, and the script continues to run.
([orphaned](https://en.wikipedia.org/wiki/Orphan_process) process is a process whose parent process has terminated, and the child continues to run. An orphaned process is adopted by init (PID 1))

### Orphaned processes

Whenever a parent process exits before its child, the child process becomes orphaned, and becomes the child of init (PID 1).

Script `create_orphan.sh` to create an orphan process
```
#!/bin/sh

sleep 20 &
sleep 10
exit
```

Run the script
```
# sh create_orphan.sh
```

From another shell run
```
# ps aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.0   4452  1584 pts/0    Ss   00:26   0:00 sh
root        50  0.0  0.0   4452  1504 pts/1    Ss   00:41   0:00 sh
root       157  0.0  0.0   4452   752 pts/0    S+   01:50   0:00 sh create_orphan.sh
root       158  0.0  0.0   4352   652 pts/0    S+   01:50   0:00 sleep 20
root       159  0.0  0.0   4352   740 pts/0    S+   01:50   0:00 sleep 10
root       160  0.0  0.0  15576  2116 pts/1    R+   01:50   0:00 ps aux
# ps -p 158 -o ppid
 PPID
  157
# ps aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.0   4452  1584 pts/0    Ss+  00:26   0:00 sh
root        50  0.0  0.0   4452  1504 pts/1    Ss   00:41   0:00 sh
root       158  0.0  0.0   4352   652 pts/0    S    01:50   0:00 sleep 20
root       162  0.0  0.0  15576  2072 pts/1    R+   01:50   0:00 ps aux
# ps -p 158 -o ppid
 PPID
    1
#
```

In the second call to `ps -p 158 -o ppid` we observe that child is adopted by init (PID 1).

### References
- [stackoverflow - create zombie in bash](https://unix.stackexchange.com/questions/217507/zombies-in-bash)
