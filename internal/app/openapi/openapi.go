// Package openapi provides handlers for fetching the OpenAPI specification for
// the service. Functionality in this package is primarily tested using API
// tests within the /tests directory of this project.
//
// OpenAPI specifications (formerly Swagger) provide a way of documenting REST
// APIs in a machine readable manner. Due to the nature and variety of services,
// it isn't always possibly to nicely encode a service, however, it is capable
// of covering all standard use cases quite well. The ability to comfortable fit
// an API into a specification can be considered a sign of good logical design;
// a rule of thumb, other factors should also be considered.
//
// They are fairly easy and quick to create and work best with JSON based
// services. A modified version of JSON schema is defined so request and
// response bodies can be specified then presented via a UI.
package openapi
