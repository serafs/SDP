package ws

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPushNotificationSubscription(t *testing.T) {
	server := NewServer()
	client := &Client{}

	t.Run("Should correctly add a session", func(t *testing.T) {
		server.addListener(client)
		require.Len(t, server.listeners, 1)
	})
}
