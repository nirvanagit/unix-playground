## Tracking logins and logouts

Files
- utmp
- wtmp
- btmp

### Analyzing them
```
[foo@bar ~]$ last -f /var/run/utmp
foo   pts/1        x.x.x.x  Sun Dec 22 05:05   still logged in
reboot   system boot  4.18.0-80.7.2.el Thu Sep 12 22:02   still running

utmp begins Thu Sep 12 22:02:56 2019
[foo@bar ~]$ sudo last -f /var/log/wtmp
foo   pts/1        x.x.x.x  Sun Dec 22 05:05   still logged in
reboot   system boot  4.18.0-80.7.2.el Thu Sep 12 22:02   still running

wtmp begins Thu Sep 12 22:02:56 2019
[foo@bar ~]$ sudo last -f /var/log/btmp

btmp begins Mon Dec  9 18:33:49 2019
[foo@bar ~]$
```
