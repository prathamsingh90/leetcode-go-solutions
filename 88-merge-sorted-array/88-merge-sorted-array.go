func merge(nums1 []int, m int, nums2 []int, n int)  {
    ind1 := m-1
    ind2 := n-1
    key := m+n-1
    
    for ind1>=0&&ind2>=0&&key>=0{
        if nums1[ind1]>=nums2[ind2]{
            nums1[key] =nums1[ind1]
            nums1[ind1]=0
            key--
            ind1--
        } else {
            nums1[key] = nums2[ind2]
            key--
            ind2--
        }
    }
    
    for ind2>=0{
        nums1[key]=nums2[ind2]
        key--
        ind2--
    }
}