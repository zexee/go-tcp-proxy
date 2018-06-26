# tcp-proxy

A small TCP proxy written in Go

## Usage

```
go get github.com/jpillora/go-tcp-proxy/cmd/tcp-proxy
```

```
$ tcp-proxy --help
Usage of tcp-proxy:
  -c: output ansi colors
  -h: output hex
  -l="localhost:9999": local address
  -n: disable nagles algorithm
  -r="localhost:80": remote address
  -v: display server actions
  -vv: display server actions and all tcp data
```

*Note: Regex match and replace*
**only works on text strings**
*and does NOT work across packet boundaries*

### Simple Example

Since HTTP runs over TCP, we can also use `tcp-proxy` as a primitive HTTP proxy:

```
$ tcp-proxy -l :9999 -r echo.jpillora.com:80
Proxying from localhost:9999 to echo.jpillora.com:80
```

### Use config file

$ tcp-proxy -c config.json

config.json is like
{
  "local": ":9999",
  "remote": "echo.jpillora.com:80",
}

#### MIT License

Copyright © 2014 Jaime Pillora <dev@jpillora.com>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
