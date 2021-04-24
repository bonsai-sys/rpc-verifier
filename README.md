Authorization workflow for (specifics ressources / api endpoint) using jwt token with external authentication server. <br />
Scope based.

Avg resp time depends on your configuration and network ping.
Best tests have been reaching <1ms.

```
BONSAI_VERIFIER_SSL		bool			= false
BONSAI_VERIFY_ADDR		string			= localhost:8972
```


```Go
r.Use(verifier.Middleware("scope; anotherscope; sscope"))
{
	r.GET("/protected_scope_route", handlers.protected_scope_route)
}

```

| 	Error Code	| Due 																						|
|:-------------:|:-----------------------------------------------------------------------------------------:|
|	400			| Invalid token length or not passed as authorization header, will try stripping Bearer. 	|
|	401			| User is not authorized with this token or /w scope. 										|
|	500			| Error encoutered while trying to authorize user. 											|



| Authorization  	|	Code	   			|
|:-----------------:|:---------------------:|
|	0				| Authorized		 	|
|	1				| Unauthorized		 	|
|	2				| Token expired			|
|	3				| Insufficent Scope		|
|	Any	(-1)		| Server Error			|