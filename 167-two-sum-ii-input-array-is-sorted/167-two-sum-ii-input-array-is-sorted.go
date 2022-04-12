func twoSum(arr []int, target int) []int {
    low := 0
    high := len(arr)-1
    ans := make([]int,2)
    
    for low<high{
        if arr[low]+arr[high]==target{
            ans[0]=low+1
            ans[1]=high+1
            break
        } else if(arr[low]+arr[high])<target{
            low++
        } else {
            high--
        }
    }
    return ans
}