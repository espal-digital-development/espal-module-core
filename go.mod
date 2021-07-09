module github.com/espal-digital-development/espal-module-core

go 1.16

replace github.com/espal-digital-development/espal-core => ../espal-core

require (
	github.com/espal-digital-development/espal-core v0.0.0-20210706103132-52ae1a2ef61c
	github.com/juju/errors v0.0.0-20200330140219-3fe23663418f
	github.com/valyala/quicktemplate v1.6.3
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
)
