// lowerChar should be passed with 65 (A in uppercase) or
// 97 (a in lowercase)
function cipher(charCode, lowerChar, k) {
  return String.fromCharCode((charCode + k - lowerChar) % 26 + lowerChar)
}
