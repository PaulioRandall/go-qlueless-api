# Changelog

All notable changes to this projects API will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Types of changes

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## Unreleased - 0.0.1

### Added

- Added `(GET) /openapi` which returns the OpenAPI specification of the API
- Added `(OPTIONS) /openapi` which handles requests for the endpoints capabilities
- Added `(GET) /changelog` which returns this changelog
- Added `(OPTIONS) /changelog` which handles requests for the endpoints capabilities
- Added `(GET) /ventures` which handles requests for Ventures
  - `ids` query parameter is a comma separated list of Venture ID's that may be used to request a subset of the data
- Added `(POST) /ventures` which handles creation of new Ventures
- Added `(PATCH) /ventures` which handles modification of existing Ventures, including deletion
- Added `(OPTIONS) /ventures` which handles requests for the endpoints capabilities
- Added `wrap` query parameter to all endpoints, except `/openapi` and `/changelog`, that will wrap the response data
  - `data` will contain the wrapped data
  - `message` contains a short summary of the response
  - `self` is the URL of the requested resource
