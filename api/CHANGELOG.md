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

## Unreleased - 2019-03-31
### Added
- Added `(GET) /openapi` which returns the OpenAPI specification of the API
- Added `(GET) /changelog` which returns this changelog
- Added `(GET) /ventures` which handles requests for `Ventures`
- Updated `(GET) /ventures` with the `id` query parameter so that specific Ventures may be requested
- Added `(POST) /ventures` which handles creation of new `Ventures`
- Added `(PUT) /ventures` which handles updating of existing `Ventures`
- Added `(HEAD) /ventures` which handles requests for `Ventures` without the body
- Added `(OPTIONS) /ventures` which handles requests for the endpoints capabilities
- THOSE BELOW ARE TO BE REMOVED
- Added `(GET) /things` which returns Things or a Thing
- Updated `(GET) /things` with the `id` query parameter so a specific Thing can be returned
- Updated `(GET) /things` to return Things or a Thing with the properties `description`, `id`, `child_ids`, `parent_ids`, `state`, `additional`, and `is_dead`
- Added `(POST) /thing` which creates a new Thing within the data store
- Updated `(POST) /thing` to accept a Thing with the properties: `description`, `child_ids`, `parent_ids`, `state`, `additional`, and `is_dead`
- Added `(PUT) /thing` which updates a Thing within the data store
- Updated `(PUT) /thing` to accept a Thing with the properties: `id`, `description`, `child_ids`, `parent_ids`, `state`, `additional`, and `is_dead`
- Added `wrap` query parameter to all endpoints, except `/openapi` and `/changelog`, that will wrap the response data to include meta information
- Updated `wrap` parameterised responses with the properties `message`, `data`, `self` and `hints`