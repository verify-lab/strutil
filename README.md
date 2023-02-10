# strutil

> Package strutil provides string helpers

![Tests](https://github.com/verify-lab/strutil/actions/workflows/tests.yml/badge.svg)

## Install

```sh
go get github.com/verify-lab/strutil
```

## Example

```go
package example

import "github.com/verify-lab/strutil"

...

fmt.Println(strutil.StripNonPrintable("example\u000a"))

// Output:
// example
```
