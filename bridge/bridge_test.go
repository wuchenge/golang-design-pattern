package bridge

func Example_common_sms() {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via SMS
}

func Example_common_email() {
	m := NewCommonMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via Emails
}

func Example_urgency_sms() {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via Email
}
