package ecs

import (
	"testing"
)

func TestGetStatus(t *testing.T) {
	newECS := ECS{
		APIUriPrefix:  "https://www.st.landui.cn",
		APISignSecret: "JMY13PagXnhl3rpiI1ht1hBBOaSF7dSOf8ktJ95zmOx19PWayRlyCtCm7UT0mghJ",
		Id:            100,
	}
	resp, code, err := newECS.GetStatus()
	if err != nil {
		t.Errorf("GetStatus failed: %s", err.Error())
	}
	t.Logf("resp: %+v code: %d", resp, code)
}
