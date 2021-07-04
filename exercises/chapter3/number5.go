package chapter3

func NewMultiSet() map[string]int {
	return map[string]int{}
}

func Insert(m map[string]int, val string) {
	if cnt, exist := m[val]; !exist {
		m[val] = 1
	} else {
		m[val] = cnt + 1
	}
}

func Erase(m map[string]int, val string) {
	if cnt, exist := m[val]; !exist {
		return
	} else if cnt > 0 {
		m[val]--
	}

	if m[val] == 0 {
		delete(m, val)
	}
}

func Count(m map[string]int, val string) int {
	if cnt, exist := m[val]; exist {
		return cnt
	}
	return 0
}

func String(m map[string]int) string {
	if len(m) == 0 {
		return "{ }"
	}

	str := "{"
	for key, val := range m {
		for i := 0; i < val; i++ {
			str += key + " "
		}
	}
	str = str[:len(str)-1] + "}"
	return str
}
