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
- Created `/orders` which returns all orders
- Created `/orders/{order_id}` which returns a specific order by ID
- Created `/batches` which returns all batches
- Created `/batches/{batch_id}` which returns a specific batch by ID
- Added `description`, `thing_id`, `parent_thing_id`, `thing_state`, and `additional` properties to both `/orders` and `/batches`
- Added `/openapi` which returns the OpenAPI specification of the API
- Added the `meta` query parameter to all `GET` endpoints except `/openapi` that will wrap the response data and include meta information within the response
- Added `message` meta information property to wrapped JSON responses that provides a summary for the response
- Added `data` meta information property to wrapped JSON responses that holds the actual response data
- Added `self` meta information property to wrapped JSON responses that holds the relative URL of the request
- Added `(POST) /order` which creates a new order within the data store