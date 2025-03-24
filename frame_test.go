package go_stomp_websocket

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Bytes(t *testing.T) {
	f := CreateFrame("CONNECTED", []string{"version:1.2", "heart-beat:1000,1000"})
	f.Body = `{"test": "json"}`
	msg := string(f.Bytes())
	expectedMsg := `["CONNECTED\nversion:1.2\nheart-beat:1000,1000\n\n{"test": "json"}\u0000"]`
	require.Equal(t, expectedMsg, msg)
}

func Test_ReadFrame(t *testing.T) {
	f := ReadFrame([]byte(`a["CONNECTED\nversion:1.2\nheart-beat:1000,1000\n\n{"test": "json"}\u0000"]`))
	expectedF := &Frame{
		Command: "CONNECTED",
		Headers: []string{"version:1.2", "heart-beat:1000,1000"},
		Body:    `{"test": "json"}`,
	}
	require.Equal(t, expectedF, f)
}

func Test_Contains(t *testing.T) {
	assertions := require.New(t)
	f := ReadFrame([]byte(`a["CONNECTED\nversion:1.2\nheart-beat:1000,1000\n\n{"test": "json"}\u0000"]`))
	version, ok := f.Contains("version")
	assertions.True(ok)
	assertions.Equal("1.2", version)
	heartBeat, ok := f.Contains("heart-beat")
	assertions.True(ok)
	assertions.Equal("1000,1000", heartBeat)
}
