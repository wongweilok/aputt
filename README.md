# APU Terminal Timetable (aputt)

## About

A small terminal application written in GO that display Asia Pacific University timetable.

## Prerequisites

* Go version 1.16.2 or above. Refer [this](https://golang.org/doc/install) page on how to install Go.

## Installation

```
$ go get -u github.com/wongweilok/aputt
```

## Third-party libraries used

1. [tcell](https://github.com/gdamore/tcell) - Licensed under [Apache License 2.0](https://github.com/gdamore/tcell/blob/master/LICENSE)
2. [tview](https://github.com/rivo/tview) - Licensed under [MIT License](https://github.com/rivo/tview/blob/master/LICENSE.txt)

## TODO List

1. Main features
   - Display timetable for current week only.
   - Color output
     * Indicate current ongoing class
     * Indicate incoming class

2. Others
   - Offline support.
   - Improve codes.

## License

aputt is released under the GPLv3 License. See [COPYING](https://github.com/wongweilok/aputt/blob/master/COPYING) for full license details.
