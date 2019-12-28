## Mounts

### mount filesystem at multiple mount points
```
# mkdir /testfs /demo
# mount /dev/loop4 /testfs
# mount /dev/loop4 /demo
# mount | grep loop4
/dev/loop4 on /home/point5 type ext4 (rw,relatime,data=ordered)
/dev/loop4 on /testfs type ext4 (rw,relatime,data=ordered)
/dev/loop4 on /demo type ext4 (rw,relatime,data=ordered)
# touch /testfs/myfile
# ls /demo
lost+found  myfile
#
```

### stacking multiple mounts on the same mount point
```
# mkdir /testfs
# mount /dev/loop4 /testfs
# touch /testfs/myfile
# mount /dev/sda1 /testfs
# mount | grep testfs
/dev/loop4 on /testfs type ext4 (rw,relatime,data=ordered)
/dev/sda1 on /testfs type ext4 (rw,relatime,data=ordered)
# touch /testfs/newfile
# ls /testfs
cni  containerd  docker  kubeadm  kubelet-plugins  log  lost+found  newfile  nfs  swap
# umount /testfs
# mount | grep testfs
/dev/loop4 on /testfs type ext4 (rw,relatime,data=ordered)
# ls /testfs
lost+found  myfile
#
```

What happened there?
We mounted /dev/sdb to /testfs and created a file myfile. Then we mounted /dev/sda1 to the same mount point, and now when we did 
`ls` we did not find /testfs/myfile. We created a /testfs/newfile and saw that it get created.

Then we unmounted /testfs, which removed the most recent mount, which was /dev/sda1. Now when we did ls /testfs we could now
again see /testfs/myfile

Uses: processes using a file descriptor from an old mount will continue to work, new processes will create files in the new
mount point.

### Mounting readonly or noexec filesystems
```
# mount /dev/loop4 /testfs
#
# mount -o noexec /dev/loop4 /demo
#
# mount | grep loop4
/dev/loop4 on /testfs type ext4 (rw,relatime,data=ordered)
/dev/loop4 on /demo type ext4 (rw,noexec,relatime,data=ordered)
# cp /bin/echo /testfs/
# cp /bin/echo /demo/

# /testfs/echo "hello world!"
hello world!

# /demo/echo "hello world!?"
sh: 174: /demo/echo: Permission denied
#
```

### Bind mounts

bind mount directory
```
# mkdir /d1
# touch /d1/x
# mkdir /d2
# mount --bind /d2 /d1
# ls /d1
x
# ls /d2
x
#
# touch /d2/y
# ls /d1
x  y
#
```

bind mount file
```
# cat > f1
Hello world!
# cat f1
Hello world!
# mount --bind f1 f2
mount: mount point f2 does not exist
# touch f2
# mount --bind f1 f2
# cat f2
Hello world!
# cat >> f2
Okay!
# cat f2
Hello world!
Okay!
# cat f1
Hello world!
Okay!
# rm f2
rm: cannot remove 'f2': Device or resource busy
# umount f2
# rm f2
# cat f1
Hello world!
Okay!
# cat f2
cat: f2: No such file or directory
#
```
