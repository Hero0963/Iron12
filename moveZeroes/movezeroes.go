package moveZeroes

func MoveZeroesMethod1(nums []int) {

	tmp := make([]int, len(nums))
	idx := 0

	for _, v := range nums {

		if v == 0 {
			continue
		}

		tmp[idx] = v
		idx++
	}

	for i, _ := range nums {
		nums[i] = tmp[i]
	}

}

func MoveZeroesMethod2(nums []int) {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

}

func MoveZeroesMethod3(nums []int) {

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for j < len(nums) {
		nums[j] = 0
		j++
	}
}
