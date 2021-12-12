package websocket_client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewWebsocketClient(t *testing.T) {
	payload := `{"intent":"pull","target":"driver","phone":""}`
	traBaseUrl := "wss://localhost:9001"
	endpointUrl := traBaseUrl + "/tra/api/v1/ws/instant/"

	got := NewWebsocketClient(endpointUrl, payload)
	fmt.Println(got)
	require.NotNilf(t, got, "Expected non nil result but received %v", got.Connection.Connection)
	require.NotNilf(t, got.Connection.Connection.OnConnected, "Expected non nil result but received %v", got.Connection.Connection.OnConnected)
	require.NotNilf(t, got.Connection.Connection.OnMessage, "Expected non nil result but received %v", got.Connection.Connection.OnMessage)
	require.NotNilf(t, got.Connection.Connection.OnError, "Expected non nil result but received %v", got.Connection.Connection.OnError)

}
