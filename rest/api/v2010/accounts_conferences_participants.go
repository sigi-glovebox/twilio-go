/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Optional parameters for the method 'CreateParticipant'
type CreateParticipantParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Whether to play a notification beep to the conference when the participant joins. Can be: `true`, `false`, `onEnter`, or `onExit`. The default value is `true`.
	Beep *string `json:"Beep,omitempty"`
	// The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that `byoc` is only meaningful when `to` is a phone number; it will otherwise be ignored. (Beta)
	Byoc *string `json:"Byoc,omitempty"`
	// The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party's phone. (Branded Calls Beta)
	CallReason *string `json:"CallReason,omitempty"`
	// The SID of the participant who is being `coached`. The participant being coached is the only participant who can hear the participant who is `coaching`.
	CallSidToCoach *string `json:"CallSidToCoach,omitempty"`
	// The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted `client:name`. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `callerId` must also be a phone number. If `to` is sip address, this value of `callerId` should be a username portion to be used to populate the From header that is passed to the SIP endpoint.
	CallerId *string `json:"CallerId,omitempty"`
	// Whether the participant is coaching another call. Can be: `true` or `false`. If not present, defaults to `false` unless `call_sid_to_coach` is defined. If `true`, `call_sid_to_coach` must be defined.
	Coaching *bool `json:"Coaching,omitempty"`
	// Whether to record the conference the participant is joining. Can be: `true`, `false`, `record-from-start`, and `do-not-record`. The default value is `false`.
	ConferenceRecord *string `json:"ConferenceRecord,omitempty"`
	// The URL we should call using the `conference_recording_status_callback_method` when the conference recording is available.
	ConferenceRecordingStatusCallback *string `json:"ConferenceRecordingStatusCallback,omitempty"`
	// The conference recording state changes that generate a call to `conference_recording_status_callback`. Can be: `in-progress`, `completed`, `failed`, and `absent`. Separate multiple values with a space, ex: `'in-progress completed failed'`
	ConferenceRecordingStatusCallbackEvent *[]string `json:"ConferenceRecordingStatusCallbackEvent,omitempty"`
	// The HTTP method we should use to call `conference_recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	ConferenceRecordingStatusCallbackMethod *string `json:"ConferenceRecordingStatusCallbackMethod,omitempty"`
	// The URL we should call using the `conference_status_callback_method` when the conference events in `conference_status_callback_event` occur. Only the value set by the first participant to join the conference is used. Subsequent `conference_status_callback` values are ignored.
	ConferenceStatusCallback *string `json:"ConferenceStatusCallback,omitempty"`
	// The conference state changes that should generate a call to `conference_status_callback`. Can be: `start`, `end`, `join`, `leave`, `mute`, `hold`, `modify`, `speaker`, and `announcement`. Separate multiple values with a space. Defaults to `start end`.
	ConferenceStatusCallbackEvent *[]string `json:"ConferenceStatusCallbackEvent,omitempty"`
	// The HTTP method we should use to call `conference_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	ConferenceStatusCallbackMethod *string `json:"ConferenceStatusCallbackMethod,omitempty"`
	// Whether to trim leading and trailing silence from your recorded conference audio files. Can be: `trim-silence` or `do-not-trim` and defaults to `trim-silence`.
	ConferenceTrim *string `json:"ConferenceTrim,omitempty"`
	// Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. Can be: `true` or `false` and defaults to `true`.
	EarlyMedia *bool `json:"EarlyMedia,omitempty"`
	// Whether to end the conference when the participant leaves. Can be: `true` or `false` and defaults to `false`.
	EndConferenceOnExit *bool `json:"EndConferenceOnExit,omitempty"`
	// The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted `client:name`. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `from` must also be a phone number. If `to` is sip address, this value of `from` should be a username portion to be used to populate the P-Asserted-Identity header that is passed to the SIP endpoint.
	From *string `json:"From,omitempty"`
	// Jitter buffer size for the connecting participant. Twilio will use this setting to apply Jitter Buffer before participant's audio is mixed into the conference. Can be: `off`, `small`, `medium`, and `large`. Default to `large`.
	JitterBufferSize *string `json:"JitterBufferSize,omitempty"`
	// A label for this participant. If one is supplied, it may subsequently be used to fetch, update or delete the participant.
	Label *string `json:"Label,omitempty"`
	// The maximum number of participants in the conference. Can be a positive integer from `2` to `250`. The default value is `250`.
	MaxParticipants *int `json:"MaxParticipants,omitempty"`
	// Whether the agent is muted in the conference. Can be `true` or `false` and the default is `false`.
	Muted *bool `json:"Muted,omitempty"`
	// Whether to record the participant and their conferences, including the time between conferences. Can be `true` or `false` and the default is `false`.
	Record *bool `json:"Record,omitempty"`
	// The recording channels for the final recording. Can be: `mono` or `dual` and the default is `mono`.
	RecordingChannels *string `json:"RecordingChannels,omitempty"`
	// The URL that we should call using the `recording_status_callback_method` when the recording status changes.
	RecordingStatusCallback *string `json:"RecordingStatusCallback,omitempty"`
	// The recording state changes that should generate a call to `recording_status_callback`. Can be: `started`, `in-progress`, `paused`, `resumed`, `stopped`, `completed`, `failed`, and `absent`. Separate multiple values with a space, ex: `'in-progress completed failed'`.
	RecordingStatusCallbackEvent *[]string `json:"RecordingStatusCallbackEvent,omitempty"`
	// The HTTP method we should use when we call `recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	RecordingStatusCallbackMethod *string `json:"RecordingStatusCallbackMethod,omitempty"`
	// The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is sent from Twilio. `both` records the audio that is received and sent by Twilio.
	RecordingTrack *string `json:"RecordingTrack,omitempty"`
	// The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:`us1`, `ie1`, `de1`, `sg1`, `br1`, `au1`, or `jp1`.
	Region *string `json:"Region,omitempty"`
	// The SIP password for authentication.
	SipAuthPassword *string `json:"SipAuthPassword,omitempty"`
	// The SIP username used for authentication.
	SipAuthUsername *string `json:"SipAuthUsername,omitempty"`
	// Whether to start the conference when the participant joins, if it has not already started. Can be: `true` or `false` and the default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
	StartConferenceOnEnter *bool `json:"StartConferenceOnEnter,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The conference state changes that should generate a call to `status_callback`. Can be: `initiated`, `ringing`, `answered`, and `completed`. Separate multiple values with a space. The default value is `completed`.
	StatusCallbackEvent *[]string `json:"StatusCallbackEvent,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` and `POST` and defaults to `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between `5` and `600`, inclusive. The default value is `60`. We always add a 5-second timeout buffer to outgoing calls, so  value of 10 would result in an actual timeout that was closer to 15 seconds.
	Timeout *int `json:"Timeout,omitempty"`
	// The phone number, SIP address, or Client identifier that received this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). SIP addresses are formatted as `sip:name@company.com`. Client identifiers are formatted `client:name`. [Custom parameters](https://www.twilio.com/docs/voice/api/conference-participant-resource#custom-parameters) may also be specified.
	To *string `json:"To,omitempty"`
	// The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file.
	WaitMethod *string `json:"WaitMethod,omitempty"`
	// The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
	WaitUrl *string `json:"WaitUrl,omitempty"`
}

func (params *CreateParticipantParams) SetPathAccountSid(PathAccountSid string) *CreateParticipantParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateParticipantParams) SetBeep(Beep string) *CreateParticipantParams {
	params.Beep = &Beep
	return params
}
func (params *CreateParticipantParams) SetByoc(Byoc string) *CreateParticipantParams {
	params.Byoc = &Byoc
	return params
}
func (params *CreateParticipantParams) SetCallReason(CallReason string) *CreateParticipantParams {
	params.CallReason = &CallReason
	return params
}
func (params *CreateParticipantParams) SetCallSidToCoach(CallSidToCoach string) *CreateParticipantParams {
	params.CallSidToCoach = &CallSidToCoach
	return params
}
func (params *CreateParticipantParams) SetCallerId(CallerId string) *CreateParticipantParams {
	params.CallerId = &CallerId
	return params
}
func (params *CreateParticipantParams) SetCoaching(Coaching bool) *CreateParticipantParams {
	params.Coaching = &Coaching
	return params
}
func (params *CreateParticipantParams) SetConferenceRecord(ConferenceRecord string) *CreateParticipantParams {
	params.ConferenceRecord = &ConferenceRecord
	return params
}
func (params *CreateParticipantParams) SetConferenceRecordingStatusCallback(ConferenceRecordingStatusCallback string) *CreateParticipantParams {
	params.ConferenceRecordingStatusCallback = &ConferenceRecordingStatusCallback
	return params
}
func (params *CreateParticipantParams) SetConferenceRecordingStatusCallbackEvent(ConferenceRecordingStatusCallbackEvent []string) *CreateParticipantParams {
	params.ConferenceRecordingStatusCallbackEvent = &ConferenceRecordingStatusCallbackEvent
	return params
}
func (params *CreateParticipantParams) SetConferenceRecordingStatusCallbackMethod(ConferenceRecordingStatusCallbackMethod string) *CreateParticipantParams {
	params.ConferenceRecordingStatusCallbackMethod = &ConferenceRecordingStatusCallbackMethod
	return params
}
func (params *CreateParticipantParams) SetConferenceStatusCallback(ConferenceStatusCallback string) *CreateParticipantParams {
	params.ConferenceStatusCallback = &ConferenceStatusCallback
	return params
}
func (params *CreateParticipantParams) SetConferenceStatusCallbackEvent(ConferenceStatusCallbackEvent []string) *CreateParticipantParams {
	params.ConferenceStatusCallbackEvent = &ConferenceStatusCallbackEvent
	return params
}
func (params *CreateParticipantParams) SetConferenceStatusCallbackMethod(ConferenceStatusCallbackMethod string) *CreateParticipantParams {
	params.ConferenceStatusCallbackMethod = &ConferenceStatusCallbackMethod
	return params
}
func (params *CreateParticipantParams) SetConferenceTrim(ConferenceTrim string) *CreateParticipantParams {
	params.ConferenceTrim = &ConferenceTrim
	return params
}
func (params *CreateParticipantParams) SetEarlyMedia(EarlyMedia bool) *CreateParticipantParams {
	params.EarlyMedia = &EarlyMedia
	return params
}
func (params *CreateParticipantParams) SetEndConferenceOnExit(EndConferenceOnExit bool) *CreateParticipantParams {
	params.EndConferenceOnExit = &EndConferenceOnExit
	return params
}
func (params *CreateParticipantParams) SetFrom(From string) *CreateParticipantParams {
	params.From = &From
	return params
}
func (params *CreateParticipantParams) SetJitterBufferSize(JitterBufferSize string) *CreateParticipantParams {
	params.JitterBufferSize = &JitterBufferSize
	return params
}
func (params *CreateParticipantParams) SetLabel(Label string) *CreateParticipantParams {
	params.Label = &Label
	return params
}
func (params *CreateParticipantParams) SetMaxParticipants(MaxParticipants int) *CreateParticipantParams {
	params.MaxParticipants = &MaxParticipants
	return params
}
func (params *CreateParticipantParams) SetMuted(Muted bool) *CreateParticipantParams {
	params.Muted = &Muted
	return params
}
func (params *CreateParticipantParams) SetRecord(Record bool) *CreateParticipantParams {
	params.Record = &Record
	return params
}
func (params *CreateParticipantParams) SetRecordingChannels(RecordingChannels string) *CreateParticipantParams {
	params.RecordingChannels = &RecordingChannels
	return params
}
func (params *CreateParticipantParams) SetRecordingStatusCallback(RecordingStatusCallback string) *CreateParticipantParams {
	params.RecordingStatusCallback = &RecordingStatusCallback
	return params
}
func (params *CreateParticipantParams) SetRecordingStatusCallbackEvent(RecordingStatusCallbackEvent []string) *CreateParticipantParams {
	params.RecordingStatusCallbackEvent = &RecordingStatusCallbackEvent
	return params
}
func (params *CreateParticipantParams) SetRecordingStatusCallbackMethod(RecordingStatusCallbackMethod string) *CreateParticipantParams {
	params.RecordingStatusCallbackMethod = &RecordingStatusCallbackMethod
	return params
}
func (params *CreateParticipantParams) SetRecordingTrack(RecordingTrack string) *CreateParticipantParams {
	params.RecordingTrack = &RecordingTrack
	return params
}
func (params *CreateParticipantParams) SetRegion(Region string) *CreateParticipantParams {
	params.Region = &Region
	return params
}
func (params *CreateParticipantParams) SetSipAuthPassword(SipAuthPassword string) *CreateParticipantParams {
	params.SipAuthPassword = &SipAuthPassword
	return params
}
func (params *CreateParticipantParams) SetSipAuthUsername(SipAuthUsername string) *CreateParticipantParams {
	params.SipAuthUsername = &SipAuthUsername
	return params
}
func (params *CreateParticipantParams) SetStartConferenceOnEnter(StartConferenceOnEnter bool) *CreateParticipantParams {
	params.StartConferenceOnEnter = &StartConferenceOnEnter
	return params
}
func (params *CreateParticipantParams) SetStatusCallback(StatusCallback string) *CreateParticipantParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateParticipantParams) SetStatusCallbackEvent(StatusCallbackEvent []string) *CreateParticipantParams {
	params.StatusCallbackEvent = &StatusCallbackEvent
	return params
}
func (params *CreateParticipantParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateParticipantParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateParticipantParams) SetTimeout(Timeout int) *CreateParticipantParams {
	params.Timeout = &Timeout
	return params
}
func (params *CreateParticipantParams) SetTo(To string) *CreateParticipantParams {
	params.To = &To
	return params
}
func (params *CreateParticipantParams) SetWaitMethod(WaitMethod string) *CreateParticipantParams {
	params.WaitMethod = &WaitMethod
	return params
}
func (params *CreateParticipantParams) SetWaitUrl(WaitUrl string) *CreateParticipantParams {
	params.WaitUrl = &WaitUrl
	return params
}

func (c *ApiService) CreateParticipant(ConferenceSid string, params *CreateParticipantParams) (*ApiV2010AccountConferenceParticipant, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Beep != nil {
		data.Set("Beep", *params.Beep)
	}
	if params != nil && params.Byoc != nil {
		data.Set("Byoc", *params.Byoc)
	}
	if params != nil && params.CallReason != nil {
		data.Set("CallReason", *params.CallReason)
	}
	if params != nil && params.CallSidToCoach != nil {
		data.Set("CallSidToCoach", *params.CallSidToCoach)
	}
	if params != nil && params.CallerId != nil {
		data.Set("CallerId", *params.CallerId)
	}
	if params != nil && params.Coaching != nil {
		data.Set("Coaching", fmt.Sprint(*params.Coaching))
	}
	if params != nil && params.ConferenceRecord != nil {
		data.Set("ConferenceRecord", *params.ConferenceRecord)
	}
	if params != nil && params.ConferenceRecordingStatusCallback != nil {
		data.Set("ConferenceRecordingStatusCallback", *params.ConferenceRecordingStatusCallback)
	}
	if params != nil && params.ConferenceRecordingStatusCallbackEvent != nil {
		for _, item := range *params.ConferenceRecordingStatusCallbackEvent {
			data.Add("ConferenceRecordingStatusCallbackEvent", item)
		}
	}
	if params != nil && params.ConferenceRecordingStatusCallbackMethod != nil {
		data.Set("ConferenceRecordingStatusCallbackMethod", *params.ConferenceRecordingStatusCallbackMethod)
	}
	if params != nil && params.ConferenceStatusCallback != nil {
		data.Set("ConferenceStatusCallback", *params.ConferenceStatusCallback)
	}
	if params != nil && params.ConferenceStatusCallbackEvent != nil {
		for _, item := range *params.ConferenceStatusCallbackEvent {
			data.Add("ConferenceStatusCallbackEvent", item)
		}
	}
	if params != nil && params.ConferenceStatusCallbackMethod != nil {
		data.Set("ConferenceStatusCallbackMethod", *params.ConferenceStatusCallbackMethod)
	}
	if params != nil && params.ConferenceTrim != nil {
		data.Set("ConferenceTrim", *params.ConferenceTrim)
	}
	if params != nil && params.EarlyMedia != nil {
		data.Set("EarlyMedia", fmt.Sprint(*params.EarlyMedia))
	}
	if params != nil && params.EndConferenceOnExit != nil {
		data.Set("EndConferenceOnExit", fmt.Sprint(*params.EndConferenceOnExit))
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.JitterBufferSize != nil {
		data.Set("JitterBufferSize", *params.JitterBufferSize)
	}
	if params != nil && params.Label != nil {
		data.Set("Label", *params.Label)
	}
	if params != nil && params.MaxParticipants != nil {
		data.Set("MaxParticipants", fmt.Sprint(*params.MaxParticipants))
	}
	if params != nil && params.Muted != nil {
		data.Set("Muted", fmt.Sprint(*params.Muted))
	}
	if params != nil && params.Record != nil {
		data.Set("Record", fmt.Sprint(*params.Record))
	}
	if params != nil && params.RecordingChannels != nil {
		data.Set("RecordingChannels", *params.RecordingChannels)
	}
	if params != nil && params.RecordingStatusCallback != nil {
		data.Set("RecordingStatusCallback", *params.RecordingStatusCallback)
	}
	if params != nil && params.RecordingStatusCallbackEvent != nil {
		for _, item := range *params.RecordingStatusCallbackEvent {
			data.Add("RecordingStatusCallbackEvent", item)
		}
	}
	if params != nil && params.RecordingStatusCallbackMethod != nil {
		data.Set("RecordingStatusCallbackMethod", *params.RecordingStatusCallbackMethod)
	}
	if params != nil && params.RecordingTrack != nil {
		data.Set("RecordingTrack", *params.RecordingTrack)
	}
	if params != nil && params.Region != nil {
		data.Set("Region", *params.Region)
	}
	if params != nil && params.SipAuthPassword != nil {
		data.Set("SipAuthPassword", *params.SipAuthPassword)
	}
	if params != nil && params.SipAuthUsername != nil {
		data.Set("SipAuthUsername", *params.SipAuthUsername)
	}
	if params != nil && params.StartConferenceOnEnter != nil {
		data.Set("StartConferenceOnEnter", fmt.Sprint(*params.StartConferenceOnEnter))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackEvent != nil {
		for _, item := range *params.StatusCallbackEvent {
			data.Add("StatusCallbackEvent", item)
		}
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.Timeout != nil {
		data.Set("Timeout", fmt.Sprint(*params.Timeout))
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.WaitMethod != nil {
		data.Set("WaitMethod", *params.WaitMethod)
	}
	if params != nil && params.WaitUrl != nil {
		data.Set("WaitUrl", *params.WaitUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountConferenceParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteParticipant'
type DeleteParticipantParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteParticipantParams) SetPathAccountSid(PathAccountSid string) *DeleteParticipantParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Kick a participant from a given conference
func (c *ApiService) DeleteParticipant(ConferenceSid string, CallSid string, params *DeleteParticipantParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchParticipant'
type FetchParticipantParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchParticipantParams) SetPathAccountSid(PathAccountSid string) *FetchParticipantParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a participant
func (c *ApiService) FetchParticipant(ConferenceSid string, CallSid string, params *FetchParticipantParams) (*ApiV2010AccountConferenceParticipant, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountConferenceParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListParticipant'
type ListParticipantParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Whether to return only participants that are muted. Can be: `true` or `false`.
	Muted *bool `json:"Muted,omitempty"`
	// Whether to return only participants that are on hold. Can be: `true` or `false`.
	Hold *bool `json:"Hold,omitempty"`
	// Whether to return only participants who are coaching another call. Can be: `true` or `false`.
	Coaching *bool `json:"Coaching,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListParticipantParams) SetPathAccountSid(PathAccountSid string) *ListParticipantParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListParticipantParams) SetMuted(Muted bool) *ListParticipantParams {
	params.Muted = &Muted
	return params
}
func (params *ListParticipantParams) SetHold(Hold bool) *ListParticipantParams {
	params.Hold = &Hold
	return params
}
func (params *ListParticipantParams) SetCoaching(Coaching bool) *ListParticipantParams {
	params.Coaching = &Coaching
	return params
}
func (params *ListParticipantParams) SetPageSize(PageSize int) *ListParticipantParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of participants belonging to the account used to make the request
func (c *ApiService) ListParticipant(ConferenceSid string, params *ListParticipantParams) (*ListParticipantResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Muted != nil {
		data.Set("Muted", fmt.Sprint(*params.Muted))
	}
	if params != nil && params.Hold != nil {
		data.Set("Hold", fmt.Sprint(*params.Hold))
	}
	if params != nil && params.Coaching != nil {
		data.Set("Coaching", fmt.Sprint(*params.Coaching))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateParticipant'
type UpdateParticipantParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The HTTP method we should use to call `announce_url`. Can be: `GET` or `POST` and defaults to `POST`.
	AnnounceMethod *string `json:"AnnounceMethod,omitempty"`
	// The URL we call using the `announce_method` for an announcement to the participant. The URL must return an MP3 file, a WAV file, or a TwiML document that contains `<Play>` or `<Say>` commands.
	AnnounceUrl *string `json:"AnnounceUrl,omitempty"`
	// Whether to play a notification beep to the conference when the participant exits. Can be: `true` or `false`.
	BeepOnExit *bool `json:"BeepOnExit,omitempty"`
	// The SID of the participant who is being `coached`. The participant being coached is the only participant who can hear the participant who is `coaching`.
	CallSidToCoach *string `json:"CallSidToCoach,omitempty"`
	// Whether the participant is coaching another call. Can be: `true` or `false`. If not present, defaults to `false` unless `call_sid_to_coach` is defined. If `true`, `call_sid_to_coach` must be defined.
	Coaching *bool `json:"Coaching,omitempty"`
	// Whether to end the conference when the participant leaves. Can be: `true` or `false` and defaults to `false`.
	EndConferenceOnExit *bool `json:"EndConferenceOnExit,omitempty"`
	// Whether the participant should be on hold. Can be: `true` or `false`. `true` puts the participant on hold, and `false` lets them rejoin the conference.
	Hold *bool `json:"Hold,omitempty"`
	// The HTTP method we should use to call `hold_url`. Can be: `GET` or `POST` and the default is `GET`.
	HoldMethod *string `json:"HoldMethod,omitempty"`
	// The URL we call using the `hold_method` for  music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains the `<Play>`, `<Say>` or `<Redirect>` commands.
	HoldUrl *string `json:"HoldUrl,omitempty"`
	// Whether the participant should be muted. Can be `true` or `false`. `true` will mute the participant, and `false` will un-mute them. Anything value other than `true` or `false` is interpreted as `false`.
	Muted *bool `json:"Muted,omitempty"`
	// The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file.
	WaitMethod *string `json:"WaitMethod,omitempty"`
	// The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
	WaitUrl *string `json:"WaitUrl,omitempty"`
}

func (params *UpdateParticipantParams) SetPathAccountSid(PathAccountSid string) *UpdateParticipantParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateParticipantParams) SetAnnounceMethod(AnnounceMethod string) *UpdateParticipantParams {
	params.AnnounceMethod = &AnnounceMethod
	return params
}
func (params *UpdateParticipantParams) SetAnnounceUrl(AnnounceUrl string) *UpdateParticipantParams {
	params.AnnounceUrl = &AnnounceUrl
	return params
}
func (params *UpdateParticipantParams) SetBeepOnExit(BeepOnExit bool) *UpdateParticipantParams {
	params.BeepOnExit = &BeepOnExit
	return params
}
func (params *UpdateParticipantParams) SetCallSidToCoach(CallSidToCoach string) *UpdateParticipantParams {
	params.CallSidToCoach = &CallSidToCoach
	return params
}
func (params *UpdateParticipantParams) SetCoaching(Coaching bool) *UpdateParticipantParams {
	params.Coaching = &Coaching
	return params
}
func (params *UpdateParticipantParams) SetEndConferenceOnExit(EndConferenceOnExit bool) *UpdateParticipantParams {
	params.EndConferenceOnExit = &EndConferenceOnExit
	return params
}
func (params *UpdateParticipantParams) SetHold(Hold bool) *UpdateParticipantParams {
	params.Hold = &Hold
	return params
}
func (params *UpdateParticipantParams) SetHoldMethod(HoldMethod string) *UpdateParticipantParams {
	params.HoldMethod = &HoldMethod
	return params
}
func (params *UpdateParticipantParams) SetHoldUrl(HoldUrl string) *UpdateParticipantParams {
	params.HoldUrl = &HoldUrl
	return params
}
func (params *UpdateParticipantParams) SetMuted(Muted bool) *UpdateParticipantParams {
	params.Muted = &Muted
	return params
}
func (params *UpdateParticipantParams) SetWaitMethod(WaitMethod string) *UpdateParticipantParams {
	params.WaitMethod = &WaitMethod
	return params
}
func (params *UpdateParticipantParams) SetWaitUrl(WaitUrl string) *UpdateParticipantParams {
	params.WaitUrl = &WaitUrl
	return params
}

// Update the properties of the participant
func (c *ApiService) UpdateParticipant(ConferenceSid string, CallSid string, params *UpdateParticipantParams) (*ApiV2010AccountConferenceParticipant, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AnnounceMethod != nil {
		data.Set("AnnounceMethod", *params.AnnounceMethod)
	}
	if params != nil && params.AnnounceUrl != nil {
		data.Set("AnnounceUrl", *params.AnnounceUrl)
	}
	if params != nil && params.BeepOnExit != nil {
		data.Set("BeepOnExit", fmt.Sprint(*params.BeepOnExit))
	}
	if params != nil && params.CallSidToCoach != nil {
		data.Set("CallSidToCoach", *params.CallSidToCoach)
	}
	if params != nil && params.Coaching != nil {
		data.Set("Coaching", fmt.Sprint(*params.Coaching))
	}
	if params != nil && params.EndConferenceOnExit != nil {
		data.Set("EndConferenceOnExit", fmt.Sprint(*params.EndConferenceOnExit))
	}
	if params != nil && params.Hold != nil {
		data.Set("Hold", fmt.Sprint(*params.Hold))
	}
	if params != nil && params.HoldMethod != nil {
		data.Set("HoldMethod", *params.HoldMethod)
	}
	if params != nil && params.HoldUrl != nil {
		data.Set("HoldUrl", *params.HoldUrl)
	}
	if params != nil && params.Muted != nil {
		data.Set("Muted", fmt.Sprint(*params.Muted))
	}
	if params != nil && params.WaitMethod != nil {
		data.Set("WaitMethod", *params.WaitMethod)
	}
	if params != nil && params.WaitUrl != nil {
		data.Set("WaitUrl", *params.WaitUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountConferenceParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
