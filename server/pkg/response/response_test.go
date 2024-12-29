package response

import (
	"errors"
	"testing"
)

func TestResponse(t *testing.T) {
	t.Log(NewResponse(200, "ok", nil))
	t.Log(NewResponse(200, "ok", map[string]any{"name": "test"}))

	t.Log(SuccessResponse(nil))
	t.Log(SuccessResponse(map[string]any{"name": "test"}))

	t.Log(ErrorResponse(-1, "error", nil))
	t.Log(ErrorResponse(-1, "error", errors.New("new error")))

	t.Log(FailedResponse(1, "failed"))
}
