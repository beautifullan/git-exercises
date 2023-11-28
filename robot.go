package exercises

type Direction int

const (
	Left Direction = iota
	Top
	Right
	Bottom
)

type Step struct {
	X int
	Y int
}

var (
	UpStep     Step = Step{X: 0, Y: 1}
	RightStep  Step = Step{X: 1, Y: 0}
	LeftStep   Step = Step{X: -1, Y: 0}
	BottomStep Step = Step{X: 0, Y: -1}
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

func (r *Robot) getStep() Step {
	switch r.Direction {
	case Left:
		return LeftStep
	case Top:
		return UpStep
	case Bottom:
		return BottomStep
	case Right:
		return RightStep
	default:
		return Step{}
	}
}

func (r *Robot) takeAction(reverse bool) {
	step := r.getStep()
	if reverse {
		r.X = r.X - step.X
		r.Y = r.Y - step.Y
	} else {
		r.X = r.X + step.X
		r.Y = r.Y + step.Y
	}
}

// 指令包括 L： 左转  R： 右转 F：向前 B：向后
func (r *Robot) Move(cmd string) {
	for _, command := range cmd {
		switch command {
		case 'L':
			r.Direction = (r.Direction + 3) % 4
		case 'R':
			r.Direction = (r.Direction + 1) % 4
		case 'F':
			r.takeAction(false)
		case 'B':
			r.takeAction(true)
		}
	}
}
