# `mirror` [![Code Coverage](https://coveralls.io/repos/github/butuzov/mirror/badge.svg?branch=main)](https://coveralls.io/github/butuzov/mirror?branch=main) [![build status](https://github.com/butuzov/mirror/actions/workflows/main.yaml/badge.svg?branch=main)]()

`mirror` help to suggest use of alternative methods, in order to avoid additional convertion of `[]byte`/`string`.

## 🇺🇦 PLEASE HELP ME 🇺🇦
Fundrise for scout drone **DJI Matrice 30T** See more details at [butuzov/README.md](https://github.com/butuzov/butuzov/)

## Examples

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
regexp.Match("foo", []byte("foobar1"))
```

</td><td>

```go
regexp.MatchString("foo", "foobar1")
```

</td></tr>
</tbody></table>

## Install

### Compile From Source

You can get `mirror` with `go install` command.

```shell
go install github.com/butuzov/mirror/cmd/mirror@latest
```

## Usage

```shell
# include test files reports
mirror --with-tests ./...
```

