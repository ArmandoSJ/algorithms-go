package two_number_sum

import "sort"

func sumTwoNumbers(integerArray []int, target int ) []int{

	sort.Ints(integerArray)
	vLeft, vRight := 0, len(integerArray) - 1
	for vLeft < vRight{
		currentSum := integerArray[vLeft] + integerArray[vRight]
		if currentSum == target{
			return []int { integerArray[vLeft], integerArray[vRight]}
		}else if currentSum < target{
			vLeft += 1
		}else{
			vRight -= 1
		}
	} 
	return []int{}
}