## File locks

```
## Locking a file
chflags uchg foobar.txt

# no opening the file in vim will show the file as RO

## unlock
chflags nouchg foobar.txt
```

