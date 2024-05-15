package ecs

import "testing"

func TestAddFirewallACL(t *testing.T) {
	acl := &AclData{
		Direction: "Inbound",
		Action:    "Allow",
		Protocol:  "TCP",
		RemoteIP:  "1.1.1.1",
		Desc:      "test",
	}
	newECS := ECS{
		APIUriPrefix:  "http://www.lxy.dev.landui.cn",
		APISignSecret: "JMY13PagXnhl3rpiI1ht1hBBOaSF7dSOf8ktJ95zmOx19PWayRlyCtCm7UT0mghJ",
		Id:            90,
		UserName:      "86326328",
	}

	err := newECS.AddFirewallACL(acl)
	if err != nil {
		t.Errorf("AddFirewallACL failed: %s", err.Error())
	}
	err = newECS.DelFirewallACL(acl)
	if err != nil {
		t.Errorf("DelFirewallACL failed: %s", err.Error())
	}
}
