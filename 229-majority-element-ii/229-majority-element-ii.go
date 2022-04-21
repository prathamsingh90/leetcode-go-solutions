func majorityElement(nums []int) []int {
    
    candidate1, candidate2 := math.MinInt, math.MinInt
    count1, count2 := 0, 0
    
    for _, num := range nums{
        if candidate1==num{
            count1++
        } else if candidate2==num{
            count2++
        } else if count1==0{
            candidate1=num
            count1=1
        } else if count2==0{
            candidate2=num
            count2=1
        } else{
            count1--
            count2--
        }
    }
    
    votes1, votes2 := 0, 0
    
    for _, num := range nums {
        if num == candidate1{
            votes1++
        } else if num == candidate2 {
            votes2++
        }
    }
    
    majority := []int{}
    if votes1>int(math.Floor(float64(len(nums)/3))){
        majority = append(majority, candidate1)
    }
    
    if votes2>int(math.Floor(float64(len(nums)/3))){
        majority = append(majority, candidate2)
    }
    
    return majority
}