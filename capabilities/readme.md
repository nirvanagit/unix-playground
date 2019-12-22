## Restrciting capabilties

[WIP]
I think runc provides an API for restricting capabilities assigned to a container, which in turn are used by docker, and kubernetes.
docker probably uses runC's [seccomp](https://github.com/opencontainers/runc/blob/7496a9682535c8ca143e981116f5b67463ad1d69/libcontainer/seccomp/seccomp_linux.go#L3)

### View capabilities
```
foo@bar:~ $ capsh --print
Current: = cap_chown,cap_dac_override,cap_fowner,cap_fsetid,cap_kill,cap_setgid,cap_setuid,cap_setpcap,cap_net_bind_service,cap_net_raw,cap_sys_chroot,cap_mknod,cap_audit_write,cap_setfcap+eip
Bounding set =cap_chown,cap_dac_override,cap_fowner,cap_fsetid,cap_kill,cap_setgid,cap_setuid,cap_setpcap,cap_net_bind_service,cap_net_raw,cap_sys_chroot,cap_mknod,cap_audit_write,cap_setfcap
Securebits: 00/0x0/1'b0
secure-noroot: no (unlocked)
secure-no-suid-fixup: no (unlocked)
secure-keep-caps: no (unlocked)
uid=0(root)
gid=0(root)
groups=
root@cfeb81ec0fab:/# grep Cap /proc/$BASHPID/status
CapInh: 00000000a80425fb
CapPrm: 00000000a80425fb
CapEff: 00000000a80425fb
CapBnd: 00000000a80425fb
CapAmb: 0000000000000000
```

Decode them using capsh
```
foo@bar:~ $ capsh --decode=00000000a80425fb
0x00000000a80425fb=cap_chown,cap_dac_override,cap_fowner,cap_fsetid,cap_kill,cap_setgid,cap_setuid,cap_setpcap,cap_net_bind_service,cap_net_raw,cap_sys_chroot,cap_mknod,cap_audit_write,cap_setfcap
foo@bar:~ $
```

### View capabilities required by a binary

For instance, to view capabilities for ping run:
```
sudo strace ping -c 1 google.com
```
