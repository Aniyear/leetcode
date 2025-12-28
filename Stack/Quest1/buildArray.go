package main

func buildArray(target []int, n int) []string {
	var ans []string
	j := 0
	for i := 1; i <= n; i++ {
		ans = append(ans, "Push")
		if target[j] == i {
			j++
			if j == len(target) {
				return ans
			}
		} else {
			ans = append(ans, "Pop")
		}
	}

	return ans
}
