# Upgrade Guide

_All `MAJOR` version bumps will have upgrade notes posted here._

[2022-05-18] 0.25.x to 0.26.x
-----------------------------
### CHANGED - Renammed ApiV2010 to Api.
ApiV2010 has now been renamed to Api. This has caused a breaking change for all endpoints located under `rest/api/2010`.

#### Send a SMS
```go
//0.2.x 
client := twilio.NewRestClient()

params := &api.CreateMessageParams{}
params.SetFrom("+15017122661")
params.SetBody("body")
params.SetTo("+15558675310")

resp, err := client.ApiV2010.CreateMessage(params)
```
```go
//1.x.xrc
client := twilio.NewRestClient()

params := &api.CreateMessageParams{}
params.SetFrom("+15017122661")
params.SetBody("body")
params.SetTo("+15558675310")

resp, err := client.Api.CreateMessage(params)
```
