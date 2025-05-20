package SequentialSearch

func soal2SeqSearch(T tabInt, N, X int) int {
	var idx, k int
	idx = 1
	k = 0
	for idx == -1 && k < N {
		if T[k] == X {
			idx = k
		}
		k = k + 1
	}
	return idx
}
