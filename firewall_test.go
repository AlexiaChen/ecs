package ecs

import "testing"

func TestAddFirewallACL(t *testing.T) {
	initLog()
	acl := &AclData{
		Direction: "Inbound",
		Action:    "Allow",
		Protocol:  "TCP",
		RemoteIP:  "1.1.1.2",
		Desc:      "test",
	}
	newECS := ECS{
		APIUriPrefix:  "http://www.lxy.dev.landui.cn",
		APISignSecret: "JMY13PagXnhl3rpiI1ht1hBBOaSF7dSOf8ktJ95zmOx19PWayRlyCtCm7UT0mghJ",
		Id:            100,
	}

	newECS.AddFirewallACL("86326328", "90", acl)
}
