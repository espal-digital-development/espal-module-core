module github.com/espal-digital-development/espal-module-core

go 1.15

replace github.com/espal-digital-development/espal-core => ../espal-core

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/espal-digital-development/espal-core v0.0.0-20201205114459-93f5a115c072
	github.com/juju/errors v0.0.0-20200330140219-3fe23663418f
	github.com/valyala/quicktemplate v1.6.3
	golang.org/x/crypto v0.0.0-20201203163018-be400aefbc4c
)
