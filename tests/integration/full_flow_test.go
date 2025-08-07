package integration

import (
	"fmt"
	"mensajeServiceIntegrationTests/componets/client"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	APIClient  *client.Client
	ServiceURL string
)

func TestMain(m *testing.M) {
	ServiceURL = os.Getenv("SERVICE_URL")
	if ServiceURL == "" {
		panic("SERVICE_URL environment variable is required")
	}

	APIClient = client.NewClient(ServiceURL)
	code := m.Run()
	os.Exit(code)
}

func TestIntegration(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	userA := fmt.Sprintf("userA-%s", timestamp)
	userB := fmt.Sprintf("userB-%s", timestamp)

	message1, err := APIClient.CreateMessage(userA, "Hola, este es mi primer mensaje!")
	time.Sleep(2 * time.Second)
	assert.NoError(t, err, "POST /messages should return 201 for first message")
	assert.NotNil(t, message1, "Response should contain message object")

	message2, err := APIClient.CreateMessage(userA, "Este es mi segundo mensaje!")
	time.Sleep(2 * time.Second)
	assert.NoError(t, err, "POST /messages should return 201 for second message")
	assert.NotNil(t, message2, "Response should contain message object")

	messages, err := APIClient.GetUserMessages(userA)
	time.Sleep(2 * time.Second)
	assert.NoError(t, err, "GET /messages should return 200")
	assert.Len(t, messages, 2, "Should return exactly 2 messages")

	messageContents := make(map[string]bool)
	for _, msg := range messages {
		messageContents[msg.Content] = true
	}
	assert.True(t, messageContents["Hola, este es mi primer mensaje!"], "Should contain first message")
	assert.True(t, messageContents["Este es mi segundo mensaje!"], "Should contain second message")
	time.Sleep(2 * time.Second)
	err = APIClient.FollowUser(userB, userA)
	assert.NoError(t, err, "POST /follows should return 201")

	time.Sleep(2 * time.Second)
	timeline, err := APIClient.GetUserTimeline(userB)
	assert.NoError(t, err, "GET /timeline should return 200")
	assert.Len(t, timeline, 2, "Timeline should contain 2 messages from followed user")

	timelineContents := make(map[string]bool)
	for _, msg := range timeline {
		timelineContents[msg.Content] = true
	}
	assert.True(t, timelineContents["Hola, este es mi primer mensaje!"], "Timeline should contain first message from userA")
	assert.True(t, timelineContents["Este es mi segundo mensaje!"], "Timeline should contain second message from userA")
}
