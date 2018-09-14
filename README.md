# JWT View
A simple utility built to help me debug JWTs

# Usage

```
Usage of ./jwtview:
  -claims
    	If specified will show the claims
  -header
    	If specified will show the header
  -signature
    	If specified will show the signature
```

## Example

```
# Token from jwt.io
$ export MYTOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

$ echo $MYTOKEN | jwtview
{"alg":"HS256","typ":"JWT"}
{"sub":"1234567890","name":"John Doe","iat":1516239022}
SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

$ echo $MYTOKEN | jwtview -claims | jq .
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
```
