func findMiddleIndex(nums []int) int {
        // sum of all elements of array
    arrSum := 0
    
    for _, num := range nums {
        arrSum += num
    }
    
    // sum of elements on left side of array
    leftSum := 0
    
    for i:=0; i<len(nums); i++ {
        arrSum -= nums[i]
        
        if arrSum == leftSum {
            return i
        } else {
            leftSum += nums[i]
        }
    }
    
    return -1

}