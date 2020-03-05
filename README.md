# twilio-go

## Documentation

The documentation for the Twilio API can be found [here][apidocs].

### Supported Go Versions

This library supports the following Go implementations:

* 1.13.8 

## Installation

To use twilio-go in your project run: 

```bash
go get github.com/twilio/twilio-go@latest
``` 



## Getting Started

Getting started with the Twilio API couldn't be easier. Create a
`Client` and you're ready to go.

### API Credentials

The `Twilio` needs your Twilio credentials. You can either pass these
directly to the constructor (see the code below) or via environment variables.

```python
from twilio.rest import Client

account = "ACXXXXXXXXXXXXXXXXX"
token = "YYYYYYYYYYYYYYYYYY"
client = Client(account, token)
```

Alternately, a `Client` constructor without these parameters will
look for `TWILIO_ACCOUNT_SID` and `TWILIO_AUTH_TOKEN` variables inside the
current environment.

We suggest storing your credentials as environment variables. Why? You'll never
have to worry about committing your credentials and accidentally posting them
somewhere public.


```python
from twilio.rest import Client
client = Client()
```

### Make a Call

```python
from twilio.rest import Client

account = "ACXXXXXXXXXXXXXXXXX"
token = "YYYYYYYYYYYYYYYYYY"
client = Client(account, token)

call = client.calls.create(to="9991231234",
                           from_="9991231234",
                           url="http://twimlets.com/holdmusic?Bucket=com.twilio.music.ambient")
print(call.sid)
```

### Send an SMS

```python
from twilio.rest import Client

account = "ACXXXXXXXXXXXXXXXXX"
token = "YYYYYYYYYYYYYYYYYY"
client = Client(account, token)

message = client.messages.create(to="+12316851234", from_="+15555555555",
                                 body="Hello there!")
```

### Handling Exceptions

```python
from twilio.rest import Client
from twilio.base.exceptions import TwilioRestException

account = "ACXXXXXXXXXXXXXXXXX"
token = "YYYYYYYYYYYYYYYYYY"
client = Client(account, token)

try:
  message = client.messages.create(to="+12316851234", from_="+15555555555",
                                   body="Hello there!")
except TwilioRestException as e:
  print(e)
```

### Generating TwiML

To control phone calls, your application needs to output [TwiML][twiml].

Use `twilio.twiml.Response` to easily create such responses.

```python
from twilio.twiml.voice_response import VoiceResponse

r = VoiceResponse()
r.say("Welcome to twilio!")
print(str(r))
```

```xml
<?xml version="1.0" encoding="utf-8"?>
<Response><Say>Welcome to twilio!</Say></Response>
```

### Docker Image

The `Dockerfile` present in this repository and its respective `twilio/twilio-python` Docker image are currently used by Twilio for testing purposes only.

### Getting help

If you need help installing or using the library, please check the [Twilio Support Help Center](https://support.twilio.com) first, and [file a support ticket](https://twilio.com/help/contact) if you don't find an answer to your question.

If you've instead found a bug in the library or would like new features added, go ahead and open issues or pull requests against this repo!

[apidocs]: https://www.twilio.com/docs/api
[twiml]: https://www.twilio.com/docs/api/twiml
