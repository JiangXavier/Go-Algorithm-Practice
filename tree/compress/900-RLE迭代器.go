type RLEIterator struct {
	num  []int
	l    int
	zhen int
	p    bool
}

func Constructor(encoding []int) RLEIterator {
	r := &RLEIterator{num: encoding, zhen: 0, l: 0, p: false}
	for i := 0; i < len(encoding); i += 2 {
		r.l += encoding[i]
	}
	return *r
}

func (this *RLEIterator) Next(n int) int {
	if this.p {
		return -1
	}
	if n > this.l {
		this.p = true
		return -1
	}
	this.l -= n
	for n > 0 {
		if this.num[this.zhen] == 0 {
			this.zhen += 2
		}
		if n <= this.num[this.zhen] {
			this.num[this.zhen] -= n
			return this.num[this.zhen+1]
		} else {
			n -= this.num[this.zhen]
			this.zhen += 2
		}
	}
	return -1
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */