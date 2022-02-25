# Changelog

## [Unreleased]

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

[Unreleased]: https://github.com/livesport-tv/util/compare/v0.2.1...master
[0.2.1]: https://github.com/livesport-tv/util/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/livesport-tv/util/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/livesport-tv/util/releases/tag/v0.1.0
