# ModelError


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **string** | Error code | [default to undefined]
**message** | **string** | Human-readable error message | [default to undefined]
**details** | **{ [key: string]: any; }** | Additional error details | [optional] [default to undefined]

## Example

```typescript
import { ModelError } from '@trackbot-app/api-client';

const instance: ModelError = {
    code,
    message,
    details,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
