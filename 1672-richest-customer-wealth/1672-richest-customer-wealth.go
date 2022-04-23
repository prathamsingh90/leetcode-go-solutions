func maximumWealth(accounts [][]int) int {
    maxWealth := math.MinInt
    
    for _, customerWealth := range accounts {
        total := 0
        for _, wealth := range customerWealth {
            total+=wealth
        }
        
        if total > maxWealth {
            maxWealth = total
        }
    }
    
    return maxWealth
}