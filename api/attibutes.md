
# Attributes

Table of attributes for all entites held and exposed by the service. The attribute keys are unique so that where two entites share an attribute key they also share the attribute meaning, its constraints, and how it should be processed; the only thing they don't share is the value.

Key | Description | Example
:--- | :--- | :---
title | Title of the entity instance | `Fix typos in legend`
description | Description of the entity instance | `The following typos have been spotted within...`
work_item_type_id | ID of a work item type | `order`
work_item_id | ID of an `order` or `batch` | `21`
parent_work_item_id | ID of a work items parent work item | `15`
tag_id | ID of the priority tag | `high`
status_id | ID of the status | `in_progress`
additional | Additional properties available for client software processing that the server doesn't process | `colour:#0000FF;last_updated:2019-03-09T15:21:08;archive_note:Some of the reported typos were just alternative spellings`