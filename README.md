# HTTP lambda Helpers

## Lamba Request

### API Gateway mapping template

```
#set($path = $input.params().path)
#set($qs = $input.params().querystring)
{
    "resourcePath": "$context.resourcePath",
    "httpMethod": "$context.httpMethod",
    "identity": {
      #foreach($key in $context.identity.keySet())
          "$key": "$context.identity.get($key)"
          #if($foreach.hasNext), #end
      #end
    },
    "authorizer": {
      #foreach($key in $context.authorizer.keySet())
          "$key": "$context.authorizer.get($key)"
          #if($foreach.hasNext), #end
      #end
    },
    "params": {
      #foreach($key in $path.keySet())
          "$key": "$path.get($key)"
          #if($foreach.hasNext), #end
      #end
    },
    "query": {
      #foreach($key in $qs.keySet())
          "$key": "$qs.get($key)"
          #if($foreach.hasNext), #end
      #end
    },
    "body": $input.json('$')
}
```

### Lambda request

```json
{
  "resourcePath": "/v1/send",
  "httpMethod": "GET",
  "identity": {
    "cognitoIdentityPoolId": "",
    "accountId": "",
    "cognitoIdentityId": "",
    "caller": "",
    "apiKey": "abc1234",
    "sourceIp": "127.0.0.1",
    "cognitoAuthenticationType": "",
    "cognitoAuthenticationProvider": "",
    "userArn": "",
    "userAgent": "PostmanRuntime/7.15.2",
    "user": ""
  },
  "params": {},
  "query": {},
  "body": {},
  "authorizer": {
    "userId": "abc1234"
  }
}
```

## Lamba Error

### API Gateway - integration error

```
#set ($errorMessageObj = $util.parseJson($input.path('$.errorMessage')))
{ "code": "$errorMessageObj.code", "message": "$errorMessageObj.public_message" }
```
