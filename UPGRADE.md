# Upgrade Guide

_All `MAJOR` version bumps will have upgrade notes posted here._

[2022-04-27] 0.2.x to 1.x.xrc
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
#### Fetch an Account
```go
//0.2.x 
client := twilio.NewRestClient()
params := &api.CreateAccountParams{}
resp, err := client.ApiV2010.FetchAccount("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
```
```go
//1.x.xrc
client := twilio.NewRestClient()
params := &api.CreateAccountParams{}
resp, err := client.Api.FetchAccount("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
```
#### Update Call Feedback
```go
//0.2.x
client := twilio.NewRestClient()

params := &api.UpdateCallFeedbackParams{}
params.SetQualityScore(1)

resp, err := client.ApiV2010.UpdateCallFeedback("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", params)
```
```go
//1.x.xrc
client := twilio.NewRestClient()

params := &api.UpdateCallFeedbackParams{}
params.SetQualityScore(1)

resp, err := client.Api.UpdateCallFeedback("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", params)
```
#### Delete a Participate From a Conference 
```go
//0.2.x
client := twilio.NewRestClient()

params := &api.DeleteParticipantParams{}

err := client.ApiV2010.DeleteParticipant("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", params)
```
```go
//1.x.xrc
client := twilio.NewRestClient()

params := &api.DeleteParticipantParams{}

err := client.Api.DeleteParticipant("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", params)
```
