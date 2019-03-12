# OpenAPI

In order to easily maintain and scale the OpenAPI specification, it has been split up. Compiling the `openapi.json` file involves running the `injector.go` script which reads `template.json` and injects files from the `/resources` folder as the template defines.

[Organisation by feature](http://www.javapractices.com/topic/TopicAction.do?Id=205) has been used as it's much easier to add, remove, and modify resources when compared with organisation by layer. Each resource has its own folder within the `/resources`. `/resources/shared` contains content referenced by multiple resources.