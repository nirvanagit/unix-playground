## Filesystem

|   Block           |  Description  |
|-------------------|-------------|
|   boot-block      |  not used by fs, contains information used to boot the OS |
|   super-block     |  contains parameter information about the fs, including size of - i-node table, logical blocks in the fs and fs in logical blocks.  |
|   i-node table    |  information about all the files |
|   data blocks     |  file data. An allocated data block can belong to only one file in the system |


## References
- [intro to linux filesystems](http://digi-cron.com:8080/filesystems.html)
