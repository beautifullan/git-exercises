package exercises

type Direction int

const (
	Left Direction = iota
	Top
	Right
	Bottom
)

// 机器人的 XY 代表了他的二维座标位置， 初始值为 (0,0)
// Direction 代表的是机器人的朝向，分别为上下左右
type Robot struct {
	X         int
	Y         int
	Direction Direction
}

func (r *Robot) Reset(direction Direction) {
	r.X = 0
	r.Y = 0
	r.Direction = direction
}

//  指令包括 L： 左转  R： 右转 F：向前 B：向后
func (r *Robot) Move(cmd string) {

}
