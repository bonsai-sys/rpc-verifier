Authorization workflow for (specifics ressources / api endpoint) using jwt token with external authentication server. <br />
Scope based.

Avg resp time depends on your configuration and network ping.
Best tests have been reaching <1ms.

```
RPC_VERIFIER_SSL			bool		= false
	// use ssl tunnel to reach distant server

RPC_VERIFIER_ADDR			string		= verify.bonsai-sys.io:8792
	// verifier end address

RPC_VERIFIER_ERROR_DEFAULT 	bool 		= false
	// Always return unauthorized as error message if true
```


```Go
r.Use(verifier.Middleware("scope; anotherscope; sscope"))
{
	r.GET("/protected_scope_route", handlers.protected_scope_route)
}
```

```Go
var Client = new(verifier.Handler)

// verifier.SetIssuer(host string)

func Auth(token Token) error {
	if token.IsBearer {
		token.StripBearer()
	}
	reply, err := Client.Authorization(token.ToString(), "scope; sscope")
	if err != nil {
		return err
	}
	// Parse(Reply) / view Reply.Parse (gin.Context)
	return nil
}
```

| 	Error Code	| Due 																						|
|:-------------:|:-----------------------------------------------------------------------------------------:|
|	401			| User is not authorized with this token or /w scope. 										|


| Authorization  	|	Code	   			|
|:-----------------:|:---------------------:|
|	0				| Authorized		 	|
|	1				| Unauthorized		 	|
|	2				| Token expired			|
|	3				| Insufficent Scope		|
|	4				| Wrong issuer			|
|	Any	(-1)		| Server Error			|