package compress

func Compress(s string) (cs string) {
	if len(s) <= 1 {
		cs = s
		return
	}
	delta := func(cnt int, lastCh rune) string {
		if cnt == 1 {
			return string(lastCh)
		} else {
			return string('0'+cnt) + string(lastCh)
		}
	}
	cnt := 1
	lastCh := rune(s[0])
	for _, ch := range s[1:] {
		if ch == lastCh {
			cnt++
		} else {
			cs += delta(cnt, lastCh)
			lastCh = ch
			cnt = 1
		}
	}
	cs += delta(cnt, lastCh)
	return
}
