func countElements(nums []int) int {
    max := math.MinInt
    min := math.MaxInt
    count := 0
    
    // Find Min and Max value in the array
    for i, _ := range nums{
        if nums[i]>max{
            max=nums[i]
        }
        
        if nums[i]<min{
            min=nums[i]
        }
    }
    
    for i, _ := range nums{
        if nums[i]<max&&nums[i]>min{
            count++
        }
    }
    
    return count
}