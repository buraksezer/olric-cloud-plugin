# olric-cloud-plugin

Discover nodes in cloud environments. Uses [hashicorp/go-discover](https://github.com/hashicorp/go-discover) under the hood.

## Install

Get the code:

```
go get -u github.com/buraksezer/olric-cloud-plugin
```

## Usage

### Load as compiled plugin

#### Build

With a properly configured Go environment:

```
go build -buildmode=plugin -o olric-cloud-plugin.so 
```

If you want to strip debug symbols from the produced binary add `-ldflags="-s -w"` to `build` command.

## Contributions

Please don't hesitate to fork the project and send a pull request or just e-mail me to ask questions and share ideas.

## License

The Apache License, Version 2.0 - see LICENSE for more details.
