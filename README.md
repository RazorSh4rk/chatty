# chatty

[![asciicast](https://asciinema.org/a/dbFvBoCfCIViyUj8WjnXeBGlB.svg)](https://asciinema.org/a/dbFvBoCfCIViyUj8WjnXeBGlB)

_Important:_ very much an early release, none of the help messages are properly
set up and there's a bunch of code that's not really needed, but it's usable otherwise

If you have go installed, you can

```bash
go build -o chatty . && sudo mv ./chatty /bin
```

Usage is explained when running, but basically
```
$ chatty -k="your api key" set
$ chatty -n="chatname" new
$ chatty talk
```

There are some other features, they are all in the help
text, just run without any arguments.