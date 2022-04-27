func sortColors(nums []int)  {
    low:=0
    mid:=0
    high:=len(nums)-1
    
    for mid=0; mid<=high;{
        if(nums[mid]==0){
            nums[low], nums[mid]=nums[mid], nums[low]
            low++
            mid++
        } else if(nums[mid]==1){
            mid++
        } else {
            nums[mid], nums[high]=nums[high],nums[mid]
            high--
        }
    }
    
}