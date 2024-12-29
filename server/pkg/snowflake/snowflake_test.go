package snowflake

import "testing"

func TestSnowflake(t *testing.T) {
	w, err := NewWorker(1)
	if err != nil {
		t.Fatal(err)
	}
	id := w.NextId()
	t.Log(id)
}
