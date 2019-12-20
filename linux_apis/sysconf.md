### Sysconf
- get configuration information at run time

### Signature
```
long sysconf(int name);
```

### Example
Get number of clock ticks per second on a system

get_clock_ticks.c

```
#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>

int main()
{
    fprintf(stdout, "No. of clock ticks per sec : %ld\n",sysconf(_SC_CLK_TCK));
    return 0;
}
```

Compile
```
gcc -o get_clock_ticks get_clock_ticks.c
```

Run
```
./get_clock_ticks
No. of clock ticks per sec : 100
```
