package main

import (
	//"log"
	"math/rand"
	"time"
)

type Bandit struct {
	Algorithm         string    `json:"algorithm"`
	Epsilon           float64   `json:"epsilon"`
	N                 int       `json:"n"`
	Counts            []int     `json:"counts"`
	Values            []float64 `json:"values"`
	ArmRewards        []float64 `json:"arm_rewards"`
	ChosenArms        []int     `json:"chosen_arms"`
	Rewards           []float64 `json:"rewards"`
	CumulativeRewards []float64 `json:"cumulative_rewards"`
}

func (b *Bandit) Initialize(algorithm string, n int, epsilon float64) {
	b.Algorithm = algorithm
	b.Epsilon = epsilon
	b.N = n

	b.ChosenArms = []int{}
	b.Rewards = []float64{}
	b.CumulativeRewards = []float64{}

	for i := 0; i < b.N; i++ {
		b.Counts = append(b.Counts, 0)
		b.Values = append(b.Values, 0)
		b.ArmRewards = append(b.ArmRewards, 0)
	}
}

func (b *Bandit) Clear() {
	b.Counts = []int{}
	b.Values = []float64{}

	b.ChosenArms = []int{}
	b.Rewards = []float64{}
	b.CumulativeRewards = []float64{}

	for i := 0; i < b.N; i++ {
		b.Counts = append(b.Counts, 0)
		b.Values = append(b.Values, 0)
		b.ArmRewards = append(b.ArmRewards, 0)
	}
}

func (b Bandit) Select_arm() int {
	rand.Seed(time.Now().UnixNano())
	p := float64(rand.Intn(100)) / 100.0

	if b.Epsilon > p {
		return rand.Intn(b.N)
	} else {
		return argmax(b.Values)
	}
}

func (b *Bandit) Update(chosen_arm int, reward float64) {
	b.Counts[chosen_arm] = b.Counts[chosen_arm] + 1
	b.ArmRewards[chosen_arm] = b.ArmRewards[chosen_arm] + reward

	n := float64(b.Counts[chosen_arm])
	value := float64(b.Values[chosen_arm])
	new_value := float64(((n-1.0)/n)*value + (1.0/n)*reward)
	b.Values[chosen_arm] = new_value

	//log.Print(b)
}

func (b *Bandit) Test_algorithm_oneshot(arms BernoulliArms) {
	chosen_arm := b.Select_arm()
	b.ChosenArms = append(b.ChosenArms, chosen_arm)
	reward := arms[chosen_arm].Draw()
	b.Rewards = append(b.Rewards, reward)
	times := len(b.CumulativeRewards)
	if times == 0 {
		b.CumulativeRewards = append(b.CumulativeRewards, reward)
	} else {
		b.CumulativeRewards = append(b.CumulativeRewards, b.CumulativeRewards[len(b.CumulativeRewards)-1]+reward)
	}
	b.Update(chosen_arm, reward)
}

func Oneshot_bandit(b *Bandit, probs []float64, epsilon float64) {
	var arms BernoulliArms

	for i, p := range probs {
		if i < b.N {
			arms = append(arms, BernoulliArm{p})
		}
	}
	b.Test_algorithm_oneshot(arms)
}

func (b *Bandit) Test_algorithm(arms BernoulliArms, num_sims int, horizon int) {
	for i := 0; i < num_sims; i++ {
		b.Clear()

		for j := 0; j < horizon; j++ {
			index := i*horizon + j
			chosen_arm := b.Select_arm()
			b.ChosenArms = append(b.ChosenArms, chosen_arm)

			reward := arms[chosen_arm].Draw()
			b.Rewards = append(b.Rewards, reward)

			if j == 0 {
				b.CumulativeRewards = append(b.CumulativeRewards, reward)
			} else {
				b.CumulativeRewards = append(b.CumulativeRewards, b.CumulativeRewards[index-1]+reward)
			}
			b.Update(chosen_arm, reward)
		}
	}
}

func Do_bandit(n_arms int, probs []float64, epsilon float64, num_sims int, horizon int) {
	var arms BernoulliArms

	for i, p := range probs {
		if i < n_arms {
			arms = append(arms, BernoulliArm{p})
		}
	}

	bandit := Bandit{}
	bandit.Initialize("EG", len(arms), epsilon)
	bandit.Test_algorithm(arms, num_sims, horizon)
	//log.Print(bandit)
}
