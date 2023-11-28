package exercises

// Palindrome 中文翻译是回文 是指一个字符串符合左右呈镜像结构
// sample:  abcdcba  日月光月日
func IsPalindrome(text string) bool {
	runes := []rune(text) //rune 类型在Go中用于表示单个Unicode字符
	half := len(runes) / 2
	for i := 0; i < half; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}
