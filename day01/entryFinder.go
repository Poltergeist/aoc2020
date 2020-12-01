package day01

func Checksum(sum int, list []int) int {
	for key, num := range list {
		if num >= 2020 {
			continue
		}
		for keyTwo, numTwo := range list {
			if numTwo >= 2020 || key == keyTwo{
				continue
			}
			if num + numTwo == 2020 {
				return num * numTwo
			}
		}
	}

	return 0
}

func ChecksumThree(sum int, list []int) int {
	for key, num := range list {
		if num >= 2020 {
			continue
		}
		for keyTwo, numTwo := range list {
			if numTwo >= 2020 || key == keyTwo{
				continue
			}
			for keyThree, numThree := range list {
				if numThree >= 2020 || key == keyThree || keyTwo == keyThree{
					continue
				}
				if num + numTwo + numThree== 2020 {
					return num * numTwo * numThree
				}
			}
		}
	}

	return 0
}
