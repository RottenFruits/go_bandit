package main

import (
	"log"
	"./bernoulliarm"
)

//type Bandit struct{
//	
//}


func foo() {
    log.Print("Hello world from foo!")
}

func bar() {
    log.Print("Hello world from bar!")
}

func main() {
	foo()
    //bar()

	var arms bernoulliarm.BernoulliArms
	probs := [2]float64{0.4, 0.8}
	for _, p := range probs{
		arms = append(arms, bernoulliarm.BernoulliArm{p})
	}

	log.Println("My favorite number is", arms)	
	log.Println("My favorite number is", arms[0])	
	log.Println("My favorite number is", arms[0].Draw())

}

