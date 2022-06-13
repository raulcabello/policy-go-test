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

func easyjson6b61f212DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *ISCSIPersistentVolumeSource) {
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
		case "chapAuthDiscovery":
			out.ChapAuthDiscovery = bool(in.Bool())
		case "chapAuthSession":
			out.ChapAuthSession = bool(in.Bool())
		case "fsType":
			out.FsType = string(in.String())
		case "initiatorName":
			out.InitiatorName = string(in.String())
		case "iqn":
			if in.IsNull() {
				in.Skip()
				out.Iqn = nil
			} else {
				if out.Iqn == nil {
					out.Iqn = new(string)
				}
				*out.Iqn = string(in.String())
			}
		case "iscsiInterface":
			out.IscsiInterface = string(in.String())
		case "lun":
			if in.IsNull() {
				in.Skip()
				out.Lun = nil
			} else {
				if out.Lun == nil {
					out.Lun = new(int32)
				}
				*out.Lun = int32(in.Int32())
			}
		case "portals":
			if in.IsNull() {
				in.Skip()
				out.Portals = nil
			} else {
				in.Delim('[')
				if out.Portals == nil {
					if !in.IsDelim(']') {
						out.Portals = make([]string, 0, 4)
					} else {
						out.Portals = []string{}
					}
				} else {
					out.Portals = (out.Portals)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Portals = append(out.Portals, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "readOnly":
			out.ReadOnly = bool(in.Bool())
		case "secretRef":
			if in.IsNull() {
				in.Skip()
				out.SecretRef = nil
			} else {
				if out.SecretRef == nil {
					out.SecretRef = new(SecretReference)
				}
				easyjson6b61f212DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.SecretRef)
			}
		case "targetPortal":
			if in.IsNull() {
				in.Skip()
				out.TargetPortal = nil
			} else {
				if out.TargetPortal == nil {
					out.TargetPortal = new(string)
				}
				*out.TargetPortal = string(in.String())
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
func easyjson6b61f212EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in ISCSIPersistentVolumeSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ChapAuthDiscovery {
		const prefix string = ",\"chapAuthDiscovery\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.ChapAuthDiscovery))
	}
	if in.ChapAuthSession {
		const prefix string = ",\"chapAuthSession\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ChapAuthSession))
	}
	if in.FsType != "" {
		const prefix string = ",\"fsType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FsType))
	}
	if in.InitiatorName != "" {
		const prefix string = ",\"initiatorName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InitiatorName))
	}
	{
		const prefix string = ",\"iqn\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Iqn == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Iqn))
		}
	}
	if in.IscsiInterface != "" {
		const prefix string = ",\"iscsiInterface\":"
		out.RawString(prefix)
		out.String(string(in.IscsiInterface))
	}
	{
		const prefix string = ",\"lun\":"
		out.RawString(prefix)
		if in.Lun == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Lun))
		}
	}
	{
		const prefix string = ",\"portals\":"
		out.RawString(prefix)
		if in.Portals == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Portals {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.ReadOnly {
		const prefix string = ",\"readOnly\":"
		out.RawString(prefix)
		out.Bool(bool(in.ReadOnly))
	}
	if in.SecretRef != nil {
		const prefix string = ",\"secretRef\":"
		out.RawString(prefix)
		easyjson6b61f212EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.SecretRef)
	}
	{
		const prefix string = ",\"targetPortal\":"
		out.RawString(prefix)
		if in.TargetPortal == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.TargetPortal))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ISCSIPersistentVolumeSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6b61f212EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ISCSIPersistentVolumeSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6b61f212EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ISCSIPersistentVolumeSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6b61f212DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ISCSIPersistentVolumeSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6b61f212DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson6b61f212DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *SecretReference) {
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
		case "name":
			out.Name = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
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
func easyjson6b61f212EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in SecretReference) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Namespace != "" {
		const prefix string = ",\"namespace\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Namespace))
	}
	out.RawByte('}')
}
