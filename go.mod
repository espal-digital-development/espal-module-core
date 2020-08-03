module github.com/espal-digital-development/espal-module-core

go 1.14

replace github.com/espal-digital-development/espal-core => ../espal-core

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/espal-digital-development/espal-core v0.0.0-20200712191757-a5faa71d3543
	github.com/juju/errors v0.0.0-20200330140219-3fe23663418f
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/valyala/quicktemplate v1.5.1
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
)
