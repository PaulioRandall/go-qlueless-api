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
- Created `/dictionaries` which returns all service dictionaries
- Added `tags` array property to `/dictionaries`; it contains all entries within the tag dictionary
- Added `title`, `description`, `tag_id`, and `additional` properties to each tag dictionary entry
- Added `statuses` array property to `/dictionaries`; it contains all entries within the status dictionary
- Added `title`, `description`, `status_id`, and `additional` properties to each status dictionary entry
- Added `work_item_types` array property to `/dictionaries`; it contains all entries within the work item type dictionary
- Added `title`, `description`, `work_item_type_id`, and `additional` properties to each work item type dictionary entry
- Created `/orders` which returns all orders
- Created `/orders/{order_id}` which returns a specific order by ID
- Created `/batches` which returns all batches
- Created `/batches/{batch_id}` which returns a specific batch by ID
- Added `title`, `description`, `work_item_id`, `parent_work_item_id`, `tag_id`, `status_id`, and `additional` properties to both `/orders` and `/batches`
- Added `/openapi` which returns the OpenAPI specification of the API
- Added the `wrap_with` query parameter to all endpoints except `/openapi` that will wrap the response data so that specific meta information can be returned for each request. The parameter accepts a dot `.` separated list of meta information properties, e.g. `wrap_with=message.self.data`
- Added `message` meta information property as an optional wrapped JSON response property. It provides a summary for the response
- Added `data` meta information property as an optional wrapped JSON response property. It holds the actual response data
- Added `self` meta information property as an optional wrapped JSON response property. It holds the relative URL of the request