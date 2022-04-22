func average(salary []int) float64 {
    min, max, sum := math.MaxInt, math.MinInt, 0
    
    for _, val := range salary {
        if val < min {
            min = val
        }
        
        if val > max {
            max = val
        }
        
        sum += val
    }
    
    sum = sum - (min+max)
    
    return (float64(sum)/float64(len(salary)-2))
}