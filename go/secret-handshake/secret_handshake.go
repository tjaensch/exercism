package secret

func Handshake(binary int) []string {
	events := []string{"wink", "double blink", "close your eyes", "jump"}

	handshake := make([]string, 0)
	if binary <= 0 {
		return handshake
	}
	//append events to handshake array for 1,10,100,1000
	for i, events := range events {
		if (1<<uint(i))&binary > 0 {
			handshake = append(handshake, events)
		}
	}

	//reverse order of events in handshake array for 10000
	if (1<<uint(4))&binary > 0 {
		for i, j := 0, len(handshake)-1; i < j; i, j = i+1, j-1 {
			handshake[i], handshake[j] = handshake[j], handshake[i]
		}
	}
	return handshake
}
