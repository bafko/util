# Livesport TV Utils

[![Go Reference](https://pkg.go.dev/badge/go.lstv.dev/util.svg)](https://pkg.go.dev/go.lstv.dev/util)
[![License: MIT](https://img.shields.io/github/license/livesport-tv/util)](https://opensource.org/licenses/MIT)
[![tests](https://github.com/livesport-tv/util/actions/workflows/tests.yml/badge.svg)](https://github.com/livesport-tv/util/actions/workflows/tests.yml)
[![Latest release](https://img.shields.io/github/v/release/livesport-tv/util?display_name=tag&sort=semver)](https://github.com/livesport-tv/util/releases)

Open source implementation of Livesport TV utilities library.

Packages provide `Formatter` and `Parser` variables to override default behavior of
format methods (`MarshalText`, `MarshalJSON`, `String`, ...)
and parse methods (`UnmarshalText`, `UnmarshalJSON`, ...).
Default value of `Formatter` is `DefaultFormatter` and default value of `Parser` is `DefaultParser`.
Also, specific verbs for `fmt` formatting are available.

## Date
```go
import "go.lstv.dev/util/date"
```

- Type `Date` represents date (year, month, day).
- Function `New` to create new date.
- Function `DateFromTime` to create date from `time.Time`.
- Type `DateFilter` to work with date intervals and filtering.

## Roman
```go
import "go.lstv.dev/util/roman"
```

- Type `Number` represents roman number.
- It supports roman numerals:
  - `M` (1000)
  - `D` (500)
  - `C` (100)
  - `L` (50)
  - `X` (10)
  - `V` (5)
  - `I` (1)
- Short and long forms are also supported and configurable.

## Sem
```go
import "go.lstv.dev/util/sem"
```

- Type `Version` represents semantic version.
- Functions:
  - `Compare`, `CompareVersion` and `CompareTag`
  - `Latest`, `LatestVersion` and `LatestTag`
  - `Parse`, `ParseVersion` and `ParseTag`
- See [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html) for more details.

## Size
```go
import "go.lstv.dev/util/size"
```

- Provides type `Size` to keep, marshal and unmarshal times-byte size values.
- It supports three JSON forms:
  - numeric form (JSON number is always in bytes)
  - string form (JSON string with or without units)
  - object form (JSON object like `{"value":1000,"unit":"MiB"}`)

## Test
```go
import "go.lstv.dev/util/test"
```

- Provides functions to test marshal and unmarshal methods:
  - `MarshalBinary` and `UnmarshalBinary`
  - `MarshalText` and `UnmarshalText`
  - `MarshalJSON` and `UnmarshalJSON`
- See [examples](./test/README.md).
