//双指针，一个读一个写

func compress(chars []byte) int {
	write, left := 0, 0
	for read, ch := range chars {
		if read == len(chars)-1 || chars[read] != chars[read+1] {
			chars[write] = ch
			write++
			num := read - left + 1
			if num != 1 {
				ji := write
				for ; num > 0; num /= 10 {
					chars[write] = '0' + byte(num%10)
					write++
				}
				s := chars[ji:write]
				for i := 0; i < len(s)/2; i++ {
					s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
				}
			}
			left = read + 1
		}
	}
	return write
}