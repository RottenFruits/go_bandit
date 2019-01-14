package main

import (
	"log"
)

func foo() {
	log.Print("Hello world from foo!")
}

func main() {
	foo()

	var arms BernoulliArms
	probs := [2]float64{0.4, 0.8}
	for _, p := range probs {
		arms = append(arms, BernoulliArm{p})
	}

	bandit := Bandit{}
	bandit.Initialize("EG", len(arms), 0.2)
	a, b, c := bandit.test_algorithm(arms, 50, 500)
	log.Print(a)
	log.Print(b)
	log.Print(c)

}
