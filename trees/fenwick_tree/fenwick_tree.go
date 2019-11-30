package tree

func update(idx, val int, bit []int) {
	idx++
	for idx <= len(bit) {
		bit[idx] += val
		idx += (idx & -idx)
	}
}

func Query(idx int, bit []int) int {
	idx++
	res := 0
	for idx > 0 {
		res += bit[idx]
		idx -= (idx & -idx)
	}
	return res
}

func Query_range(left, right int, bit []int) int {
	return Query(right, bit) - Query(left-1, bit)
}
