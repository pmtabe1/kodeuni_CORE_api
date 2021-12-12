package webrtc_client

type IWebrtcClient interface {
}

type WebrtcClient struct {
}

func New() *WebrtcClient {

	return &WebrtcClient{}
}
