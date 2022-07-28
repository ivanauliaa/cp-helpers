package golang

// lowerChar should be passed with 65 (A in uppercase) or
// 97 (a in lowercase)
func cipher(char *byte, lowerChar byte, k int32) string {
	return string((*char+byte(k)-lowerChar)%26 + lowerChar)
}
