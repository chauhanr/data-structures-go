package algorithms

func FindLongestSubsequence(s []int, length []int){
	for k :=1; k< len(s); k++{
		length[k] = 1
		for i :=0; i<k; i++ {
			if s[i] < s[k]{
				length[k] = maxValue(length[k], length[i]+1)
			}
		}
	}
}

func maxValue(a, b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}


