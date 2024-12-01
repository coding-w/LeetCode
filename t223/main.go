package main

func main() {

}

// https://leetcode.cn/problems/rectangle-area/
// 根据覆盖的特性，找到覆盖的矩形
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	s := (ax1-ax2)*(ay1-ay2) + (bx1-bx2)*(by1-by2)
	if ax1 > bx2 || ay1 > by2 || ax2 < bx1 || ay2 < by1 {
		return s
	}
	x1 := max(ax1, bx1)
	y1 := max(ay1, by1)
	x2 := min(ax2, bx2)
	y2 := min(ay2, by2)
	return s - (x2-x1)*(y2-y1)
}
