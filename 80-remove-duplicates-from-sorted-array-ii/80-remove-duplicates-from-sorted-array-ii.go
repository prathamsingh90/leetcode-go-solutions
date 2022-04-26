func removeDuplicates(nums []int) int {
    if len(nums)==0||len(nums)==1{
        return len(nums)
    }
    
    orig:=1
    count:=1
    
    for i:=1;i<=len(nums)-1;i++{
        if(nums[i]==nums[i-1]){
            count++
            if(count<=2){
                nums[orig]=nums[i]
                orig++
            }
        } else {
            count=1
            nums[orig]=nums[i]
            orig++
        }
    }    
    return orig
}