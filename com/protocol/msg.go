package protocol

var ProtocolMap = make(map[string][]interface{}, 0)

func init() {
	var protocol = make([]interface{}, 0, 100)
	protocol = append(protocol,
		(*LoginReq)(nil),
		(*LoginAck)(nil),
		(*LogoutReq)(nil),
		(*LogoutAck)(nil),
		(*GetGameRoomTypeListReq)(nil),
		(*GetGameRoomTypeListAck)(nil),
	)

	ProtocolMap["protocol"] = protocol
}
