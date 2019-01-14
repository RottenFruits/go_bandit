package bernoulliarm

import (
	"math/rand"
	"time"
)

type BernoulliArm struct{
	P float64
}

type BernoulliArms []BernoulliArm

func (b BernoulliArm) Draw() int{
	var p float64

	rand.Seed(time.Now().UnixNano())
	p = float64(rand.Intn(100)) / 100.0

	if p > b.P {
		return 0
	}else{
		return 1
	}
}