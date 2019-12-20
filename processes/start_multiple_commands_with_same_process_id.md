### Why?
This is helpful if you want to manipulate multiple processes using pgid

Create two files first.sh and second.sh with the following content
```
#!/bin/sh

sleep 120
```

Run
```
[foo@bar processes]$ setsid sh -c 'sh ./first.sh & sh ./second.sh'
```

Check
```
[foo@bar processes]$ ps -e -o pid,ppid,pgid,command | grep 1157
 1157     1  1157 sh -c sh ./first.sh & sh ./second.sh
 1158  1157  1157 sh ./first.sh
 1159  1157  1157 sh ./second.sh
 1160  1159  1157 sleep 120
 1161  1158  1157 sleep 120
 1455  8482  1454 grep --color=auto 1157
[foo@bar processes]$
```

Killing all processes under 1157 (parent process)
```
[foo@bar processes]$ kill -TERM -1157
[foo@bar processes]$ ps -e -o pid,ppid,pgid,command | grep 1157
 1794  8482  1793 grep --color=auto 1157
[foo@bar processes]$
```
