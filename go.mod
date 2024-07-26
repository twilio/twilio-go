module github.com/ghostmonitor/twilio-go

go 1.16

require (
	github.com/beevik/etree v1.1.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/mock v1.6.0
	github.com/kr/text v0.2.0 // indirect
	github.com/localtunnel/go-localtunnel v0.0.0-20170326223115-8a804488f275
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

retract (
    v1.22.0 // Contains known bug
)
