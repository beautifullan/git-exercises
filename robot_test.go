package exercises

import (
	"testing"
)

func TestRobotMove(t *testing.T) {
	robot := Robot{}
	robot.Direction = Top
	robot.Move("LFFRB")

	if robot.X != -2 && robot.Y != -1 {
		t.Fail()
	}

	robot.Reset(Bottom)
	robot.Move("LFFRB")
	if robot.X != 2 && robot.Y != 1 {
		t.Fail()
	}
}
