package twiml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/twiml"
)

func TestVoice_ZeroValueVerbList(t *testing.T) {
	var verbs []twiml.Element
	resp, _ := twiml.Voice(verbs)
	assert.Equal(t,
		`<?xml version="1.0" encoding="UTF-8"?><Response/>`,
		resp)
}

func TestVoice_EmptyVerbList(t *testing.T) {
	verbs := []twiml.Element{}
	resp, _ := twiml.Voice(verbs)
	assert.Equal(t,
		`<?xml version="1.0" encoding="UTF-8"?><Response/>`,
		resp)
}

func TestVoice_SayVerb(t *testing.T) {
	say := twiml.VoiceSay{
		Message: "Hello World!",
	}
	verbs := []twiml.Element{say}
	resp, _ := twiml.Messages(verbs)
	assert.Equal(t,
		`<?xml version="1.0" encoding="UTF-8"?><Response><Say>Hello World!</Say></Response>`,
		resp)
}

func TestVoice_SayVerbWithParams(t *testing.T) {
	say := twiml.VoiceSay{
		Message:  "Hello World!",
		Voice:    "alice",
		Loop:     "1",
		Language: "en",
	}
	verbs := []twiml.Element{say}
	resp, _ := twiml.Messages(verbs)
	assert.Contains(t, resp, `voice="alice"`)
	assert.Contains(t, resp, `loop="1"`)
	assert.Contains(t, resp, `language="en"`)
}

func TestVoice_SayVerbWithOptionalAttributes(t *testing.T) {
	optAttrs := map[string]string{"PlaybackURL": "https://demo.twilio.com"}
	say := twiml.VoiceSay{
		Message:            "Hello World!",
		OptionalAttributes: optAttrs,
	}
	verbs := []twiml.Element{say}
	resp, _ := twiml.Messages(verbs)
	assert.Contains(t, resp, `playbackURL="https://demo.twilio.com"`)
}

func TestVoice_GatherVerbWithNestedPlayVerb(t *testing.T) {
	play := twiml.VoicePlay{
		Url: "https://demo.twilio.com",
	}
	gather := twiml.VoiceGather{
		InnerElements: []twiml.Element{play},
	}
	verbs := []twiml.Element{gather}
	resp, _ := twiml.Messages(verbs)
	assert.Equal(t,
		`<?xml version="1.0" encoding="UTF-8"?><Response><Gather><Play>https://demo.twilio.com</Play></Gather></Response>`,
		resp)
}
