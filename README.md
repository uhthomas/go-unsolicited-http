# Go Unsolicited HTTP

This repository is a reproduction for an issue in the Go HTTP package which
where data is incorrectly being shared across requests. I don't know if the
problem is the client, or the server.

## How to use

Run the server:

```sh
❯ go run github.com/uhthomas/gounsolicitedhttp/cmd/server@main
```

Run the client:

```sh
❯ go run github.com/uhthomas/gounsolicitedhttp/cmd/client@main
```

## The problem

### Quick Failure

Server:

```sh
attempt 0
attempt 1
2024/07/26 16:15:47 do: Head "http://localhost:8080": net/http: HTTP/1.x transport connection broken: malformed HTTP status code "html>"
exit status 1
```

Client:

```sh
attempt 0
attempt 1
2024/07/26 16:15:47 do: Head "http://localhost:8080": net/http: HTTP/1.x transport connection broken: malformed HTTP status code "html>"
exit status 1
```

### Slow Failure

Server:

```sh
1: 54 bytes written, err=<nil>
2: 0 bytes written, err=readfrom tcp [::1]:8080->[::1]:48336: write tcp [::1]:8080->[::1]:48336: write: broken pipe
3: 54 bytes written, err=<nil>
4: 54 bytes written, err=<nil>
5: 54 bytes written, err=<nil>
6: 54 bytes written, err=<nil>
7: 54 bytes written, err=<nil>
8: 54 bytes written, err=<nil>
9: 54 bytes written, err=<nil>
10: 0 bytes written, err=readfrom tcp [::1]:8080->[::1]:33742: write tcp [::1]:8080->[::1]:33742: write: broken pipe
```

Client:

```sh
attempt 0
2024/07/26 16:17:17 Unsolicited response received on idle HTTP channel starting with "<!DOCTYPE html>\n<html>\n  <head>\n    <title>Thomas</tit"; err=<nil>
attempt 1
2024/07/26 16:17:17 Unsolicited response received on idle HTTP channel starting with "<!DOCTYPE html>\n<html>\n  <head>\n    <title>Thomas</tit"; err=<nil>
attempt 2
2024/07/26 16:17:17 Unsolicited response received on idle HTTP channel starting with "<!DOCTYPE html>\n<html>\n  <head>\n    <title>Thomas</tit"; err=<nil>
attempt 3
2024/07/26 16:17:17 Unsolicited response received on idle HTTP channel starting with "<!DOCTYPE html>\n<html>\n  <head>\n    <title>Thomas</tit"; err=<nil>
attempt 4
2024/07/26 16:17:17 Unsolicited response received on idle HTTP channel starting with "<!DOCTYPE html>\n<html>\n  <head>\n    <title>Thomas</tit"; err=<nil>
attempt 5
2024/07/26 16:17:17 Unsolicited response received on idle HTTP channel starting with "<!DOCTYPE html>\n<html>\n  <head>\n    <title>Thomas</tit"; err=<nil>
attempt 6
attempt 7
2024/07/26 16:17:18 do: Head "http://localhost:8080": net/http: HTTP/1.x transport connection broken: malformed HTTP status code "html>"
exit status 1
``
