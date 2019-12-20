### Start process as child of pid 1

sleep.sh
```
#!/bin/sh

sleep 120
```

Run
```
setsid sh ./sleep.sh
```

Check ppid
```
[foo@bar processes]$ ps -p 30518 -o pid,ppid,pgid,command
  PID  PPID  PGID COMMAND
30518     1 30518 sh ./sleep.sh
```

PPID of this process is set to 1, and not the pid of the login shell
