package model

type SDP string

type MessageType string

var (
	MessageTypeNotifyClientID MessageType = "notify-client-id"
	MessageTypeSDPOffer       MessageType = "offer"
	MessageTypeSDPAnswer      MessageType = "answer"
	MessageTypeICECandidate   MessageType = "ice-candidate"
	MessageTypeNewClient      MessageType = "new-client"
	MessageTypeLeaveClient    MessageType = "leave-client"
)

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload"`
}
