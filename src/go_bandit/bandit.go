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

type banditResults struct {
	Chosen_arms        []int     `json:chosen_arms`
	Rewards            []float64 `json:rewards`
	Cumulative_rewards []float64 `json:cumulative_rewards`
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

func (b *Bandit) Update(chosen_arm int, reward float64) {
	b.Counts[chosen_arm] = b.Counts[chosen_arm] + 1
	n := float64(b.Counts[chosen_arm])
	value := float64(b.Values[chosen_arm])
	new_value := float64(((n-1.0)/n)*value + (1.0/n)*reward)
	b.Values[chosen_arm] = new_value
}

func (b Bandit) Test_algorithm(arms BernoulliArms, num_sims int, horizon int) banditResults {
	var bandit_results banditResults

	for i := 0; i < num_sims; i++ {
		b.Clear()

		for j := 0; j < horizon; j++ {
			index := i*horizon + j
			chosen_arm := b.Select_arm()
			bandit_results.Chosen_arms = append(bandit_results.Chosen_arms, chosen_arm)

			reward := arms[chosen_arm].Draw()
			bandit_results.Rewards = append(bandit_results.Rewards, reward)

			if j == 0 {
				bandit_results.Cumulative_rewards = append(bandit_results.Cumulative_rewards, reward)
			} else {
				bandit_results.Cumulative_rewards = append(bandit_results.Cumulative_rewards, bandit_results.Cumulative_rewards[index-1]+reward)
			}
			b.Update(chosen_arm, reward)
		}
	}

	return bandit_results
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
	bandit_results := bandit.Test_algorithm(arms, num_sims, horizon)
	log.Print(bandit_results)
}
