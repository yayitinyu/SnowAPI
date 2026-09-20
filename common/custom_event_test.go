package common

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomEventRender(t *testing.T) {
	recorder := httptest.NewRecorder()
	event := CustomEvent{Data: "data: hello"}

	require.NoError(t, event.Render(recorder))
	assert.Equal(t, "text/event-stream", recorder.Header().Get("Content-Type"))
	assert.Equal(t, "no-cache", recorder.Header().Get("Cache-Control"))
	assert.Equal(t, "data: hello\n\n", recorder.Body.String())
}

func TestCustomEventWriteContentTypePreservesCacheControl(t *testing.T) {
	recorder := httptest.NewRecorder()
	recorder.Header().Set("Cache-Control", "private")

	CustomEvent{}.WriteContentType(recorder)

	assert.Equal(t, "private", recorder.Header().Get("Cache-Control"))
}
