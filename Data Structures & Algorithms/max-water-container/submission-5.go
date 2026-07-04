func maxArea(heights []int) int {
    
	
	lp := 0
	rp := len(heights)-1
	maxAmt  := 0

	for lp < rp {

		minheight := min(heights[lp],heights[rp])
		width := rp -lp

		area := width * minheight

		if area > maxAmt{
			maxAmt=area
		}
		if heights[lp] < heights[rp] {
			lp++
		}else{
			rp--
		}
	}
   return maxAmt
}
