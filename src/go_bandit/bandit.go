package main

import (
	"log"
	"math/rand"
	"time"
)

type Bandit struct {
	Algorithm string
	Epsilon   float64
	N         int
	Counts    []int
	Values    []float64
}

func (b *Bandit) Initialize(algorithm string, n int, epsilon float64) {
	b.Algorithm = algorithm
	b.Epsilon = epsilon
	b.N = n

	for i := 0; i < b.N; i++ {
		b.Counts = append(b.Counts, 0)
		b.Values = append(b.Values, 0)
	}
}

func (b *Bandit) Clear() {
	b.Counts = []int{}
	b.Values = []float64{}

	for i := 0; i < b.N; i++ {
		b.Counts = append(b.Counts, 0)
		b.Values = append(b.Values, 0)
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

func (b *Bandit) update(chosen_arm int, reward float64) {
	b.Counts[chosen_arm] = b.Counts[chosen_arm] + 1
	n := float64(b.Counts[chosen_arm])
	value := float64(b.Values[chosen_arm])
	new_value := float64(((n-1.0)/n)*value + (1.0/n)*reward)
	b.Values[chosen_arm] = new_value
}

func (b Bandit) test_algorithm(arms BernoulliArms, num_sims int, horizon int) ([]int, []float64, []float64) {
	chosen_arms := []int{}
	rewards := []float64{}
	cumulative_rewards := []float64{}

	for i := 0; i < num_sims; i++ {
		b.Clear()

		for j := 0; j < horizon; j++ {
			index := i*horizon + j
			chosen_arm := b.Select_arm()
			chosen_arms = append(chosen_arms, chosen_arm)

			reward := arms[chosen_arm].Draw()
			rewards = append(rewards, reward)

			if j == 0 {
				cumulative_rewards = append(cumulative_rewards, reward)
			} else {
				cumulative_rewards = append(cumulative_rewards, cumulative_rewards[index-1]+reward)
			}
			b.update(chosen_arm, reward)
		}
	}

	return chosen_arms, rewards, cumulative_rewards
}

func Do_bandit(n_arms int, probs []float64, epsilon float64, num_sims int, horizon int) {
	var arms BernoulliArms

	for _, p := range probs {
		arms = append(arms, BernoulliArm{p})
	}

	bandit := Bandit{}
	bandit.Initialize("EG", len(arms), epsilon)

	a, b, c := bandit.test_algorithm(arms, num_sims, horizon)
	var res = banditResults{a, b, c}
	log.Print(res)
}
