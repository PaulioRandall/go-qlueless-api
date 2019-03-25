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

## Unreleased - 2019-03-13
### Added
- Added `(GET) /openapi` which returns the OpenAPI specification of the API
- Added `(GET) /changelog` which returns this changelog
- Added `(GET) /things` which returns all Things
- Added `(GET) /things/{id}` which returns a specific Thing
- Updated `(GET) /things` and `(GET) /things/{id}` to return Things or a Thing with the properties `description`, `id`, `child_ids`, `state`, `additional`, and `is_dead`
- Added `(POST) /thing` which creates a new Thing within the data store
- Updated `(POST) /thing` to accept a Thing with the properties: `description`, `child_ids`, `state`, `additional`, and `is_dead`
- Added the `wrap` query parameter to all `GET` endpoints, except `(GET) /openapi` and `(GET) /changelog`, that will wrap the response data to include meta information
- Updated `wrap` parameterised responses with the properties `message`, `data` and `self`