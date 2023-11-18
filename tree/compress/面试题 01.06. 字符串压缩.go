// strings.Builder   高效创建字符串
// WriteByte 方法写byte字符
// WriteString 方法写进字符串
// strconv.Itoa 整数转换字符类型
// Len() 方法获得长度 ， String() 方法获得其字符类型转换
func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}
	var j strings.Builder
	cur := S[0]
	curlen := 1
	for i := 1; i < len(S); i++ {
		if S[i] == cur {
			curlen++
		} else if S[i] != cur {
			j.WriteByte(cur)
			j.WriteString(strconv.Itoa(curlen))
			cur = S[i]
			curlen = 1
		}
	}
	j.WriteByte(cur)
	j.WriteString(strconv.Itoa(curlen))
	if j.Len() >= len(S) {
		return S
	}
	return j.String()
}