# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- `dpdocs` will now accept the eServices Membership report as an input file as an alternative to CAPWATCH.
- The Barcode Attendance log now sorts names.
- Join Date and Rank Date are now a part of the Domain Model.
- New report: schedule of cadet feedback sessions by month.

### Changed
- Refactored LaTeX report generation into a common package.

### Deprecated

### Removed

### Fixed

### Security

## [1.0.0]

### Changed
- Output file name now defaults to the log date
- `dpdocs` will ask for the CAPWATCH password if it will be refreshing the cache instead of throwing an error

### Removed
- `--password` flag removed from command-line interface

### Security
- Instead of typing the eServices password on the command line in the clear, a no-echo scan approach is now used

## [0.1.0]

### Added
- Fetching membership data from CAPWATCH
- Parsing squadron Table of Organization data from a YAML config file
- Creating Barcode Attendance Logs from CAPWATCH data and YAML config

[unreleased]: https://github.com/ut080/bcs-portal/compare/v0.1.0...HEAD
[1.0.0]: https://github.com/ut080/bcs-portal/compare/v0.1.0...v1.0.0
[0.1.0]: https://github.com/ut080/bcs-portal/releases/tag/v0.1.0