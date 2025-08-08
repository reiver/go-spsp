# go-spsp

Package **spsp** implements the **Simple Payment Setup Protocol** (**SPSP**), for the Go programming language.

The **Simple Payment Setup Protocol** (**SPSP**) is used by the [Interledger Protocol](https://interledger.org/), [Open Payments](https://openpayments.dev/), [Web Monetization](https://webmonetization.org/).

An HTTPS request (with the appropriate HTTP `Accept` header) to a resolve **Payment Pointer** can return a **Simple Payment Setup Protocol** (**SPSP**) _response_.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-spsp

[![GoDoc](https://godoc.org/github.com/reiver/go-spsp?status.svg)](https://godoc.org/github.com/reiver/go-spsp)

## Import

To import package **spsp** use `import` code like the following:
```
import "github.com/reiver/go-spsp"
```

## Installation

To install package **spsp** do the following:
```
GOPROXY=direct go get github.com/reiver/go-spsp
```

## Author

Package **spsp** was written by [Charles Iliya Krempeaux](http://reiver.link)

## See Also

* https://github.com/reiver/go-ilpaddr
* https://github.com/reiver/go-pymtptr
* https://github.com/reiver/go-spsp
