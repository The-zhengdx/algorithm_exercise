package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	if l == 1 {
		return (n == 1 && flowerbed[0] == 0) || n == 0
	}
	if l == 2 {
		return n == 0 || (n == 1 && flowerbed[0] == 0 && flowerbed[1] == 0)
	}
	for i := 0; n > 0 && i < l; i++ {
		if flowerbed[i] == 0 {
			if i > 0 && i < l-1 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			} else if i == 0 && flowerbed[1] == 0 {
				flowerbed[i] = 1
				n--
			} else if i == l-1 && flowerbed[l-2] == 0 {
				flowerbed[i] = 1
				n--
			}
		}
	}
	return n == 0
}
