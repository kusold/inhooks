package services

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/didil/inhooks/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMessageDecoderFromHttp_OK(t *testing.T) {
	flowID := "flow-1"
	jsonPayload := []byte(`{"id":"1234","status":"complete"}`)
	b := bytes.NewBuffer(jsonPayload)
	rawQuery := "x=abc&yz=this%20is%20ok"
	r := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/ingest/%s?%s", flowID, rawQuery), b)

	r.Header = http.Header{
		"header-1": []string{"abc"},
		"header-2": []string{"def"},
	}

	flow := &models.Flow{
		ID: flowID,
	}

	d := NewMessageDecoder()
	message, err := d.FromHttp(flow, r)
	assert.NoError(t, err)

	_, err = uuid.Parse(message.ID)
	assert.NoError(t, err)

	assert.Equal(t, flowID, message.FlowID)
	assert.Equal(t, rawQuery, message.RawQuery)
	assert.Equal(t, r.Header, message.HttpHeaders)
	assert.Equal(t, jsonPayload, message.Payload)
}
