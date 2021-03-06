# Changelog

## [Unreleased]

## [0.8.0] - 2022-05-14
### Added
- Panic handling for `test.*` functions.
- Field `Before` and `After` to `test.Case*` types.
- New error checkers:
  - `test.AnyError`
  - `test.ErrorHasPrefix`
  - `test.ErrorHasSuffix`
  - `test.ErrorMatch`

## [0.7.1] - 2022-04-27
### Fixed
- Package `test` allows to test also pointer types.

## [0.7.0] - 2022-04-25
### Added
- Format/parse and `fmt.Formatter` support for `uu` package.
- Function `uu.RandomID`.

### Fixed
- Error messages.

## [0.6.0] - 2022-04-19
### Added
- Package `uu` with UUID support.

## [0.5.0] - 2022-03-23
### Added
- Format `roman.FormatLowerCase`.
- Methods `sem.Ver.Core` and `sem.Ver.IsZero`.
- Implementation of `fmt.Formatter` for:
  - `date.Date`
  - `roman.Number`
  - `sem.Ver`
- Package `test` for enhanced testing.

### Changed
- Implementation of tests.

### Fixed
- Method `Date.MarshalBinary` error.

## [0.4.0] - 2022-03-16
### Added
- Package `constraint` with generic support and helpers.
- Function `size.Bytes` with generic support.
- Methods `sem.Ver.NextMajor`, `sem.Ver.NextMinor` and `sem.Ver.NextPatch`.

### Changed
- Type `Rule` doc comments (all packages).
- Used `any` instead of `interface{}`.
- Generic support:
  - Error types.
  - Function `size.New`
  - All parser functions now accepts `constraint.ParserInput`.
- Updated:
  - Go to 1.18
  - `github.com/stretchr/testify` to `1.7.1`

### Removed
- Methods of `size.Size`:
  - `BytesInt`
  - `BytesUint`
  - `BytesInt32`
  - `BytesUint32`
  - `BytesInt64`
  - `BytesUint64`
  - `BytesFloat32`
  - `BytesFloat64`

## [0.3.0] - 2022-02-28
### Added
- Type `Rule` and `DefaultParser` parameter of same type (all packages).

### Changed
- Renamed all unmarshal global functions to parse (all packages).
  - Format/Parse and Formatter/Parser is better choice than Format/UnmarshalText.
  - Variables `UnmarshalText` renamed to `Parser`.
  - Functions `DefaultUnmarshalText` renamed to `DefaultParser`.
- Published under-laying errors (all packages).
- Unified `input` and `data` names.
  - Renamed `MaxTextLength` to `MaxInputLength`.

### Removed
- Global unmarshal flag variables (all packages).

### Fixed
- Doc comments.

## [0.2.1] - 2022-02-25
### Fixed
- Added doc comments for `size.ObjectKeyValue` and `size.ObjectKeyUnit`.
- Method `sem.Ver.Compare` and function `sem.DefaultComparePreRelease` comparison of release and pre-release.

## [0.2.0] - 2022-02-24
### Added
- Introduced new constants `size.ObjectKeyValue` and `size.ObjectKeyUnit`.
- Missing doc comments.

### Changed
- Renamed `size.NewSize` to `size.New`.
- Optimized `size.unmarshalJSONObject` implementation.

### Removed
- Variable `sem.Zero`.
  - Use `sem.Ver{}` instead.
- Name of `date.filterNo.Contains` receiver.

## [0.1.0] - 2022-02-22
### Added
- First release of util.

[Unreleased]: https://github.com/livesport-tv/util/compare/v0.8.0...master
[0.8.0]: https://github.com/livesport-tv/util/compare/v0.7.1...v0.8.0
[0.7.1]: https://github.com/livesport-tv/util/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/livesport-tv/util/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/livesport-tv/util/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/livesport-tv/util/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/livesport-tv/util/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/livesport-tv/util/compare/v0.2.1...v0.3.0
[0.2.1]: https://github.com/livesport-tv/util/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/livesport-tv/util/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/livesport-tv/util/releases/tag/v0.1.0
