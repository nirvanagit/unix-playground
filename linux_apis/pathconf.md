### Pathconf
- get configuration for a file

### Signature
```
long fpathconf(int fd, int name);
long pathconf(const char *path, int name);
```

### Example
Get the max number of bytes for which space is available in a terminal input queue.

max_bytes_in_terminal_input_queue.c

```
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main( void )
{
  long value;

  // pathconf("filename", paramater) examples of parameter can be found on [gnu manual](https://www.gnu.org/software/libc/manual/html_node/Pathconf.html)
  value = pathconf( "/dev/con1", _PC_MAX_INPUT );
  printf( "Input buffer size is %ld bytes\n", value );
  return EXIT_SUCCESS;
}
```

Compile
```
gcc -o max_bytes_in_terminal_input_queue max_bytes_in_terminal_input_queue.c
```

Run
```
./max_bytes_in_terminal_input_queue
Input buffer size is 255 bytes
```
