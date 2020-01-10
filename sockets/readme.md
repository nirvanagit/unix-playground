## Sockets

System calls associated to sockets

-  [__socket()__](http://man7.org/linux/man-pages/man2/socket.2.html) => creates a new socket

```
int socket(int domain, int type, int protocol);
```

-   [__bind()__](http://man7.org/linux/man-pages/man2/bind.2.html) => binds a socket to an address, this enables clients to find it a connect to it.

```
int bind(int sockfd, const struct sockaddr *addr,
    socklen_t addrlen);

```

-   [__listen()__](http://man7.org/linux/man-pages/man2/listen.2.html) => allows a stream socket to access incoming connections from other sockets.

```
int listen(int sockfd, int backlog);
```

-   [__accept()__](http://man7.org/linux/man-pages/man2/accept.2.html) => accepts a connection from peer

```
int accept4(int sockfd, struct sockaddr *addr,
    socklen_t *addrlen, int flags);
```

-   [__connect()__](http://man7.org/linux/man-pages/man2/connect.2.html) => establishes connection with another socket.

```
int connect(int sockfd, const struct sockaddr *addr,
    socklen_t addrlen);
```
