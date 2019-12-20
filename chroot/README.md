## chroot

## Steps
```
## step 1
mkdir -p jail/usr/lib/system
mkdir jail/bin

## step 2
### copy relevant programs
cp /bin/sh jail/bin/
cp /bin/ls jail/bin
cp /bin/pwd jail/bin/

## step 3
### copy required libraries
sudo cp -r /usr/lib/* jail/usr/lib/

### might also require this
cp -r /lib64 jail/

## step 4
### enter chroot
sudo chroot jail /bin/sh
$ sudo chroot jail /bin/sh
sh-3.2# ls
bin     usr
sh-3.2# pwd
/
sh-3.2#
```
