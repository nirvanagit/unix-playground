##
### Have a running container

Get container process id
```
CONTAINER_PID=$(docker inspect -f '{{.State.Pid}}' <container_id>)
```

Get capabilities assigned to this process
```
grep Cap /proc/$CONTAINER_PID/status
CapInh: 00000000a80425fb
CapPrm: 00000000a80425fb
CapEff: 00000000a80425fb
CapBnd: 00000000a80425fb
CapAmb: 0000000000000000
```

Decode the above capabilities using capsh
```
foo@bar:~ $ capsh --decode=00000000a80425fb
0x00000000a80425fb=cap_chown,cap_dac_override,cap_fowner,cap_fsetid,cap_kill,cap_setgid,cap_setuid,cap_setpcap,cap_net_bind_service,cap_net_raw,cap_sys_chroot,cap_mknod,cap_audit_write,cap_setfcap
foo@bar:~ $
```
