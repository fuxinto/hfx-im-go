package pb

import "google.golang.org/protobuf/proto"

// NewFrom new packet from a header
func NewFrom(packType PackType, body proto.Message) ([]byte, error) {

	data, err := proto.Marshal(body)
	if err != nil {
		return nil, err
	}
	pkt := &Pack{Type: packType, Body: data}
	return proto.Marshal(pkt)
}
