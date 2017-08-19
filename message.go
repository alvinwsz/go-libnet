package libnet

import (
	"encoding/gob"
	"encoding/json"
	"encoding/xml"

	"github.com/alvinwsz/glog"
)

// Convert to bytes message.
func Bytes(v []byte) Encoder {
	return func(buffer *OutBuffer) error {
		buffer.WriteBytes(v)
		return nil
	}
}

// Convert to string message.
func String(v string) Encoder {
	return func(buffer *OutBuffer) error {
		buffer.WriteString(v)
		return nil
	}
}

// Create a json message.
func Json(v interface{}) Encoder {
	return func(buffer *OutBuffer) error {
		//glog.Info("before: ", v);
		err := json.NewEncoder(buffer).Encode(v)
		glog.Info(">>>>>>TX(Json): ", string(buffer.Data))
		return err
	}
}

// Create a gob message.
func Gob(v interface{}) Encoder {
	return func(buffer *OutBuffer) error {
		return gob.NewEncoder(buffer).Encode(v)
	}
}

// Create a xml message.
func Xml(v interface{}) Encoder {
	return func(buffer *OutBuffer) error {
		return xml.NewEncoder(buffer).Encode(v)
	}
}

//增加id是为了在broadcast时不向具有同一id的session发消息
//这是为了防止broadcast时给自己发消息

// Convert to bytes message.
func (session *Session) Bytes(v []byte) Encoder {
	return func(buffer *OutBuffer) error {
		buffer.id = session.Id()
		buffer.WriteBytes(v)
		return nil
	}
}

// Convert to string message.
func (session *Session) String(v string) Encoder {
	return func(buffer *OutBuffer) error {
		buffer.id = session.Id()
		buffer.WriteString(v)
		return nil
	}
}

// Create a json message.
func (session *Session) Json(v interface{}) Encoder {
	return func(buffer *OutBuffer) error {
		buffer.id = session.Id()
		return json.NewEncoder(buffer).Encode(v)
	}
}

// Create a gob message.
func (session *Session) Gob(v interface{}) Encoder {
	return func(buffer *OutBuffer) error {
		buffer.id = session.Id()
		return gob.NewEncoder(buffer).Encode(v)
	}
}

// Create a xml message.
func (session *Session) Xml(v interface{}) Encoder {
	return func(buffer *OutBuffer) error {
		buffer.id = session.Id()
		return xml.NewEncoder(buffer).Encode(v)
	}
}

//==============================================
// Decoder
func JsonDecoder(v interface{}) Decoder {
	return func(buffer *InBuffer) error {
		return json.NewDecoder(buffer).Decode(v)
	}
}

func GobDecoder(v interface{}) Decoder {
	return func(buffer *InBuffer) error {
		return gob.NewDecoder(buffer).Decode(v)
	}
}

func XmlDecoder(v interface{}) Decoder {
	return func(buffer *InBuffer) error {
		return xml.NewDecoder(buffer).Decode(v)
	}
}
