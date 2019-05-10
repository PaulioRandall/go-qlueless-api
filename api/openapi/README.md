# OpenAPI

In order to easily maintain and scale the OpenAPI specification, it has been split up. Compiling the `openapi.json` file involves running one of the scripts within `/scripts` which reads `template.json` and injects files from the `/resources` folder as the template defines.

[Organisation by feature](http://www.javapractices.com/topic/TopicAction.do?Id=205) has been used as it's much easier to add, remove, and modify resources when compared with organisation by layer. Each resource has its own folder within the `/resources`. `/resources/shared` contains content referenced by multiple resources.

## Format and Example

The following will inject the contents of `{resources-path}/shared/parameters.json` into the template and indent it by 2 tabs:

**Resources: {resources-path}/shared/parameters.json**

```json
"wrap": {
  "name": "wrap",
  "in": "query",
  "description": "If present wraps the response and adds meta information.",
  "required": false,
  "schema": {
    "type": "string"
  }
}
```

**Template:**

```json
{
  "parameters": {
    {{- "\n"}}{{ .Inject "/shared/parameters.json" 2}}
  }
}
```

**Output:**

```json
{
  "parameters": {
    "wrap": {
      "name": "wrap",
      "in": "query",
      "description": "If present wraps the response and adds meta information.",
      "required": false,
      "schema": {
        "type": "string"
      }
    }
  }
}
```