package secret

var steps = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(input uint) []string {
	var handshake []string
	for i := 0; i < 4; i++ {
		if input&(1<<uint(i)) == (1 << uint(i)) {
			if input < 16 {
				handshake = append(handshake, steps[i])
				continue
			}
			handshake = append([]string{steps[i]}, handshake...)
		}
	}
	return handshake
}