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

func easyjson32242f71DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *AzureFileVolumeSource) {
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
		case "readOnly":
			out.ReadOnly = bool(in.Bool())
		case "secretName":
			if in.IsNull() {
				in.Skip()
				out.SecretName = nil
			} else {
				if out.SecretName == nil {
					out.SecretName = new(string)
				}
				*out.SecretName = string(in.String())
			}
		case "shareName":
			if in.IsNull() {
				in.Skip()
				out.ShareName = nil
			} else {
				if out.ShareName == nil {
					out.ShareName = new(string)
				}
				*out.ShareName = string(in.String())
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
func easyjson32242f71EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in AzureFileVolumeSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ReadOnly {
		const prefix string = ",\"readOnly\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.ReadOnly))
	}
	{
		const prefix string = ",\"secretName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.SecretName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.SecretName))
		}
	}
	{
		const prefix string = ",\"shareName\":"
		out.RawString(prefix)
		if in.ShareName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ShareName))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AzureFileVolumeSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson32242f71EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AzureFileVolumeSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson32242f71EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AzureFileVolumeSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson32242f71DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AzureFileVolumeSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson32242f71DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
