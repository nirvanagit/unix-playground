# Creating filesystems and mounting them

Presetup
```
docker run -it --privileged=true ubuntu:14.04 sh

# mount
overlay on / type overlay (rw,relatime,lowerdir=/var/lib/docker/overlay2/l/KCQFK5GL6IRFEQNG2O5SOPPX7P:/var/lib/docker/overlay2/l/GCTWRQ66HSHDLU2LSCFHKXS4O3:/var/lib/docker/overlay2/l/PFJREKOSQGPLWHPM2VYHHX7EMJ:/var/lib/docker/overlay2/l/754R6HCR7EGN3HXUG4LM4HRB4C,upperdir=/var/lib/docker/overlay2/abd7830dfff2317e7e70c8cd596230da607da9643f96709fb5e97cb5e30da8a5/diff,workdir=/var/lib/docker/overlay2/abd7830dfff2317e7e70c8cd596230da607da9643f96709fb5e97cb5e30da8a5/work)
proc on /proc type proc (rw,nosuid,nodev,noexec,relatime)
tmpfs on /dev type tmpfs (rw,nosuid,size=65536k,mode=755)
devpts on /dev/pts type devpts (rw,nosuid,noexec,relatime,gid=5,mode=620,ptmxmode=666)
sysfs on /sys type sysfs (rw,nosuid,nodev,noexec,relatime)
tmpfs on /sys/fs/cgroup type tmpfs (rw,nosuid,nodev,noexec,relatime,mode=755)
cpuset on /sys/fs/cgroup/cpuset type cgroup (rw,nosuid,nodev,noexec,relatime,cpuset)
cpu on /sys/fs/cgroup/cpu type cgroup (rw,nosuid,nodev,noexec,relatime,cpu)
cpuacct on /sys/fs/cgroup/cpuacct type cgroup (rw,nosuid,nodev,noexec,relatime,cpuacct)
blkio on /sys/fs/cgroup/blkio type cgroup (rw,nosuid,nodev,noexec,relatime,blkio)
memory on /sys/fs/cgroup/memory type cgroup (rw,nosuid,nodev,noexec,relatime,memory)
devices on /sys/fs/cgroup/devices type cgroup (rw,nosuid,nodev,noexec,relatime,devices)
freezer on /sys/fs/cgroup/freezer type cgroup (rw,nosuid,nodev,noexec,relatime,freezer)
net_cls on /sys/fs/cgroup/net_cls type cgroup (rw,nosuid,nodev,noexec,relatime,net_cls)
perf_event on /sys/fs/cgroup/perf_event type cgroup (rw,nosuid,nodev,noexec,relatime,perf_event)
net_prio on /sys/fs/cgroup/net_prio type cgroup (rw,nosuid,nodev,noexec,relatime,net_prio)
hugetlb on /sys/fs/cgroup/hugetlb type cgroup (rw,nosuid,nodev,noexec,relatime,hugetlb)
pids on /sys/fs/cgroup/pids type cgroup (rw,nosuid,nodev,noexec,relatime,pids)
cgroup on /sys/fs/cgroup/systemd type cgroup (rw,nosuid,nodev,noexec,relatime,name=systemd)
mqueue on /dev/mqueue type mqueue (rw,nosuid,nodev,noexec,relatime)
/dev/sda1 on /etc/resolv.conf type ext4 (rw,relatime,data=ordered)
/dev/sda1 on /etc/hostname type ext4 (rw,relatime,data=ordered)
/dev/sda1 on /etc/hosts type ext4 (rw,relatime,data=ordered)
shm on /dev/shm type tmpfs (rw,nosuid,nodev,noexec,relatime,size=65536k)
devpts on /dev/console type devpts (rw,nosuid,noexec,relatime,gid=5,mode=620,ptmxmode=666)

# create mount points
# pwd
/home
# mkdir point1 point2 point3 point4
#
```

## Commands used

### mount

### mkfs

### dd

### mknod

### losetup

## Mounting 

### tmpfs filesystem
```
# mount -t tmpfs mount1 point1
# mount | grep mount1
mount1 on /home/point1 type tmpfs (rw,relatime)
# ls point1/
#
```

### proc filesystem
```
# mount -t proc mount2 point2
# mount | grep mount2
mount2 on /home/point2 type proc (rw,relatime)
# ls point2/
1          bus        consoles  diskstats    fb           iomem     kcore      kpagecgroup  locks    mounts  pagetypeinfo  slabinfo  sys            timer_list  vmallocinfo
36         cgroups    cpuinfo   dma          filesystems  ioports   key-users  kpagecount   meminfo  mpt     partitions    softirqs  sysrq-trigger  tty         vmstat
acpi       cmdline    crypto    driver       fs           irq       keys       kpageflags   misc     mtrr    sched_debug   stat      sysvipc        uptime      zoneinfo
buddyinfo  config.gz  devices   execdomains  interrupts   kallsyms  kmsg       loadavg      modules  net     self          swaps     thread-self    version
#
```

### devpts filesystem
```
# mount -t devpts mount3 point3
# ls point3/
ptmx
#
# mount | grep mount3
mount3 on /home/point3 type devpts (rw,relatime,mode=600,ptmxmode=000)
#
```

### mqueue filesystem
```
# mount -t mqueue mount4 point4
# ls point4/
# mount | grep point4
mount4 on /home/point4 type mqueue (rw,relatime)
#
```

### mount sysfs filesystem
```
# mkdir point6
# mount -t sysfs mount6 point6
# mount | grep point6
mount6 on /home/point6 type sysfs (rw,relatime)
#
# ls point6
block  bus  class  dev  devices  firmware  fs  hypervisor  kernel  module  power
#
```

### mount ext3/ext4 filesystem

This is a special type and a simple mount command would not work here. To create a ext3 filesystem, we need a device,
for our purpose we will create a virtual device by using Linux's [loop device](http://man7.org/linux/man-pages/man4/loop.4.html)

```
# dd if=/dev/zero of=virtual_hard_drive bs=1M count=1024
1024+0 records in
1024+0 records out
1073741824 bytes (1.1 GB) copied, 1.6688 s, 643 MB/s
# losetup /dev/loop4 virtual_hard_drive
# mkfs -t ext4 /dev/loop4
mke2fs 1.42.9 (4-Feb-2014)
Discarding device blocks: done
Filesystem label=
OS type: Linux
Block size=4096 (log=2)
Fragment size=4096 (log=2)
Stride=0 blocks, Stripe width=0 blocks
65536 inodes, 262144 blocks
13107 blocks (5.00%) reserved for the super user
First data block=0
Maximum filesystem blocks=268435456
8 block groups
32768 blocks per group, 32768 fragments per group
8192 inodes per group
Superblock backups stored on blocks:
        32768, 98304, 163840, 229376

Allocating group tables: done
Writing inode tables: done
Creating journal (8192 blocks): done
Writing superblocks and filesystem accounting information: done

# mkdir point5
# mount /dev/loop4 point5
# mount | grep point5
/dev/loop4 on /home/point5 type ext4 (rw,relatime,data=ordered)
# ls /home/point5/
lost+found
#
```

## References

- [virtual-hard-drive](https://superuser.com/questions/686064/how-to-attach-a-virtual-hard-drive-to-occupy-dev-sdb)
- [using mkfs](https://www.thegeekdiary.com/how-to-create-and-mount-filesystems-in-linux/)
- [loop device](http://man7.org/linux/man-pages/man4/loop.4.html)
