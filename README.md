# Go Kadena

GO client library for interacting with a Kadena chainweb node.
<br/>
It also contain API wrapper functions for interacting with Pact's development server. The API for interacting with Pact server is the same as described in [pact documentation](https://pact-language.readthedocs.io/en/stable/pact-reference.html?highlight=%2Fsend#rest-api)

## Install

```bash
   go install github.com/jfamousket/go-kadena@latest
```

## Functions

### Hashing

Create a `blake2b` hash

```go
import (
    "github.com/jfamousket/go-kadena/helpers"
)

CreateBlake2Hash(<cmd_string>) string
```

### Pact API functions

This functions are same as described in [pact-api-documentation](https://pact-language.readthedocs.io/en/stable/pact-reference.html?highlight=%2Fsend#rest-api)

```go
import (
    "github.com/jfamousket/go-kadena/fetch"
)

Send(<valid_cmd_object>, <api_host>) SendResponse
Local(<valid_cmd_object>, <api_host>) LocalResponse
Listen(<request_key>, <api_host>) ListenResponse
Poll(<array_of_request_keys>, <api_host>) PollResponse
```

## TODOS

- [ ] Key generation and manipulation functions
- [x] Create wallet Kadena blockchain
- [ ] Other expected features for blockchain library
