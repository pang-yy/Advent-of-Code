package utils

func BytesToInt(byteArray []byte) int {
	n := 0
	for _, b := range byteArray {
		if b >= '0' && b <= '9' {
		//if b != ' ' {
			n += int(b - '0')
			n *= 10
		}
	}
	n /= 10
	return n
}
