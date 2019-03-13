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
- Created `/orders` which returns a dummy order
- Created `/batches` which returns a list of dummy batches
- Added `title`, `description`, `work_item_id`, `parent_work_item_id`, `tag_id`, `status_id`, and `additional` properties to both `/orders` and `/batches`
- Added `/openapi` which returns the OpenAPI specification of the API
- Added the `wrap` query parameter to all endpoints except `/openapi` that will wrap the response data so meta information can be obtained for each request
- Added `message` text property to the top level of all `wrap`ped JSON responses to provide a summary for each response
- Added `data` object property to the top level of all `wrap`ped JSON responses to hold the actual response data