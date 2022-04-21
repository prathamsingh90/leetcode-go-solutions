func majorityElement(nums []int) []int {
    
    elemCount := map[int]int{}
    list := []int{}
    
    for _, num := range nums {
        
       elemCount[num] += 1
    }
    
    for num, count := range elemCount{
        if count > int(math.Floor(float64(len(nums)/3))) {
            list = append(list, num)
        }
    }
    
    return list
}