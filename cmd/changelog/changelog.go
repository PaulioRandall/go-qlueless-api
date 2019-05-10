// Package changelog provides handlers for fetching the API changelog for the
// API. Functionality in this package is primarily tested using API tests
// within the /tests directory of this project.
//
// Changelogs for an API provide a way for consumers to see what changes have
// been made to it over time. Particularly useful for client developers during
// preventative maintenance and upgrades to check for fixes and new features
// so they can decide whether to make changes their end.
//
// Regardless of the actual versioning approach used, versioning is key to a
// good changelog. Sensible client developers will document the version they
// are targetting and expect to easily find all changes since then in the
// API changelog.
//
// I've used a changelog here to show the changes to the API itself, not the
// source code.
//
// Keep a Changelog: https://keepachangelog.com/en/1.0.0/
package changelog
