// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson38da1406DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *SessionAffinityConfig) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "clientIP":
			if in.IsNull() {
				in.Skip()
				out.ClientIP = nil
			} else {
				if out.ClientIP == nil {
					out.ClientIP = new(ClientIPConfig)
				}
				(*out.ClientIP).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson38da1406EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in SessionAffinityConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ClientIP != nil {
		const prefix string = ",\"clientIP\":"
		first = false
		out.RawString(prefix[1:])
		(*in.ClientIP).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SessionAffinityConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson38da1406EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SessionAffinityConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson38da1406EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SessionAffinityConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson38da1406DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SessionAffinityConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson38da1406DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
