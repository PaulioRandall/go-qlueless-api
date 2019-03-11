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

## Unreleased - 2019-03-10
### Added
- Added `message` text property to the top level of all JSON responses to provide a summary for each response
- Added `data` object property to the top level of all JSON responses to group all specific response data
- Created `/dictionaires` endpoint which returns all service dictionaires
- Added `tags` array property to `/dictionaires#data`; it contains all entries within the tag dictionary
- Added `title`, `description`, `tag_id`, and `additional` properties to each tag dictionary entry
- Added `statuses` array property to `/dictionaires#data`; it contains all entries within the status dictionary
- Added `title`, `description`, `status_id`, and `additional` properties to each status dictionary entry
- Added `work_item_types` array property to `/dictionaires#data`; it contains all entries within the work item type dictionary
- Added `title`, `description`, `work_item_type_id`, and `additional` properties to each work item type dictionary entry
- Created `/orders` endpoint which returns a dummy order
- Created `/batches` endpoint which returns a list of dummy batches
- Added `title`, `description`, `work_item_id`, `parent_work_item_id`, `tag_id`, `status_id`, and `additional` properties to both `/orders#data` and `/batches#data`
- Added `/openapi` endpoint which returns the OpenAPI specification of the API