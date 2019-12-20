## User Management

In a docker environment, we will create two groups and multiple users and see how the `/home` directory looks like

```
docker run -it centos:latest sh

sh-4.4# groupadd admin
sh-4.4# useradd ironman -G admin

sh-4.4# groupadd play
sh-4.4# useradd foo -G play
sh-4.4# useradd bar -G play

sh-4.4# ls /home/
aaeron  bar  foo
sh-4.4#
```
