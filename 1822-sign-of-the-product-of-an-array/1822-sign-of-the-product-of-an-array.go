func arraySign(nums []int) int {
    negNum := 0
    
    for _, num := range nums{
        if num <0{
            negNum++
        } else if num==0{
            return 0
        }
    }
    
    if negNum%2==0{
        return 1
    }
    
    return -1
}