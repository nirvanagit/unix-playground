## Debugging using [delve](https://github.com/go-delve/delve)

```
# cd into directory containing main.go

➜  program git:(master) ✗ dlv debug
Type 'help' for list of commands.
# list breakpoints
(dlv) breakpoints
Breakpoint runtime-fatal-throw at 0x10282c0 for runtime.fatalthrow() /usr/local/go/src/runtime/panic.go:820 (0)
Breakpoint unrecovered-panic at 0x1028330 for runtime.fatalpanic() /usr/local/go/src/runtime/panic.go:847 (0)
        print runtime.curg._panic.arg

# set breakpoint
(dlv) break main.go:26
Breakpoint 1 set at 0x1058273 for main.main() ./main.go:26
(dlv) continue
> main.main() ./main.go:26 (hits goroutine(1):1 total:1) (PC: 0x1058273)
    21:         var father Person
    22:         father.Name = "goo"
    23:         father.Age = 1
    24:
    25:         var family []Person
=>  26:         family = append(family, me, wife, mother, father)
    27:
    28:         birthdayToday(&me)
    29: }
    30:
    31: func birthdayToday(person *Person) {

(dlv) print father
main.Person {Name: "goo", Age: 1}
(dlv) print mother
main.Person {Name: "eoo", Age: 1}
(dlv) print me
main.Person {Name: "foo", Age: 0}
(dlv) print wife
main.Person {Name: "bar", Age: 0}
(dlv) print family
[]main.Person len: 0, cap: 0, nil
(dlv) print father
main.Person {Name: "goo", Age: 1}

# change values of non string properties
(dlv) set father.Age = 2
(dlv) print father
main.Person {Name: "goo", Age: 2}

# use n to go to the next line and use that as the breakpoint
(dlv) n
> main.main() ./main.go:28 (PC: 0x10583a1)
    23:         father.Age = 1
    24:
    25:         var family []Person
    26:         family = append(family, me, wife, mother, father)
    27:
=>  28:         birthdayToday(&me)
    29: }
    30:
    31: func birthdayToday(person *Person) {
    32:         person.Age = person.Age + 1
    33: }
(dlv) print family
[]main.Person len: 4, cap: 4, [
        {Name: "foo", Age: 0},
        {Name: "bar", Age: 0},
        {Name: "eoo", Age: 1},
        {Name: "goo", Age: 2},
]
(dlv) quit
➜  program git:(master) ✗
```

### References
- [rakyll](https://rakyll.org/coredumps/)
- [gdb dashboard](https://github.com/cyrus-and/gdb-dashboard)
- [debugging go apps using delve](https://www.youtube.com/watch?v=qFf2PRSfBlQ)
