func majorityElement(nums []int) int {
    
    count := 0
    candidate:=math.MinInt
    
    for i, _ := range nums {
        if count==0{
            count=1
            candidate=nums[i]
        } else if candidate==nums[i]{
            count++
        } else {
            count--
        }
    }
    
    votes:=0
    
    for i, _ := range nums {
        if nums[i]==candidate{
            votes++
        }
    }
    
    if votes >= int(math.Floor(float64(len(nums)/2))){
        return candidate
    }
    
    return -1
}