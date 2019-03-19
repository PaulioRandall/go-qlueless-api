# Changelog
All notable changes to this project will be documented in this file.

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
- Created `/things` which returns all things
- Created `/things/{id}` which returns a specific thing by ID
- Added `description`, `id`, `parent_id`, `state`, and `additional` properties to `/things`
- Added `/openapi` which returns the OpenAPI specification of the API
- Added the `meta` query parameter to all `GET` endpoints except `/openapi` that will wrap the response data and include meta information within the response
- Added `message` meta information property to wrapped JSON responses that provides a summary for the response
- Added `data` meta information property to wrapped JSON responses that holds the actual response data
- Added `self` meta information property to wrapped JSON responses that holds the relative URL of the request
- Added `(POST) /thing` which creates a new thing within the data store