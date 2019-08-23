package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	testCases := []struct {
		status  bool
		message string
	}{
		{
			status:  true,
			message: "status should be true",
		},
		{
			status:  false,
			message: "status should be false",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.message, func(t *testing.T) {
			result := Message(tc.status, tc.message)

			require.Equal(t, tc.status, result["status"])
			require.Equal(t, tc.message, result["message"])
		})
	}
}

func TestRespond(t *testing.T) {
	data := map[string]interface{}{"status": true, "message": "testing"}
	w := httptest.NewRecorder()

	Respond(w, data)

	require.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
