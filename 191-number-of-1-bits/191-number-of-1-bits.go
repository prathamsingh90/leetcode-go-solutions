func hammingWeight(num uint32) int {
    
    numOneBits := bits.OnesCount32(num)
    return numOneBits
}