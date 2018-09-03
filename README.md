# xmlfmt

A streaming XML formatter, written in Go (golang).

:warning: This tool is not production ready.

### Install

**Binaries**

[![Releases](https://img.shields.io/github/release/jpillora/xmlfmt.svg)](https://github.com/jpillora/xmlfmt/releases) [![Releases](https://img.shields.io/github/downloads/jpillora/xmlfmt/total.svg)](https://github.com/jpillora/xmlfmt/releases)

See [the latest release](https://github.com/jpillora/xmlfmt/releases/latest) or download and install it now with `curl https://i.jpillora.com/xmlfmt! | bash`

**Source**

```sh
$ go get -v github.com/jpillora/xmlfmt
```

### Usage

```
$ xmlfmt --help

  Usage: xmlfmt [options] [file]

  path to xml [file] (default -)

  Options:
  --write, -w      write over xml file with formatted
  --max-width, -m  max width of the file in characters (default unlimited)
  --help, -h
  --version, -v

  Version:
    X.Y.Z

  Read more:
    github.com/jpillora/xmlfmt
```

#### MIT License

Copyright © 2018 Jaime Pillora &lt;dev@jpillora.com&gt;

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
