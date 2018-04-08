package mobile_notificaion_service

//`json:"omitempty"`
type AndroidNotification struct {
	Id         string
	TypeEnum   int
	Title      string
	Body       string
	MeUserId   int
	PeerUserId int
	PostId     int
	GroupId    int
	NotifyId   int
	Sound      string
	Vibrate    string
	ChatKey    string
}
