package jwt

import "testing"

func TestGenerateToken(t *testing.T) {
	id := int64(1)
	sessionID := "test"
	token, err := GenerateToken(id, sessionID)
	if err != nil {
		t.Errorf("GenerateToken failed: %v", err)
	}
	t.Logf("GenerateToken success: %s", token)
}

func TestCheck(t *testing.T) {
	id := int64(1)
	sessionID := "test"
	token, err := GenerateToken(id, sessionID)
	if err != nil {
		t.Errorf("GenerateToken failed: %v", err)
	}
	claims, ok := Check(token)
	if !ok {
		t.Errorf("Check failed: %v", err)
	}
	t.Logf("Check success: %v", claims)
}
