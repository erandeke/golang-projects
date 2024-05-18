package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting go routine demo....")

	evalNinjas := []string{"Tom", "Dick", "Harry"}

	for _, ninjaBhau := range evalNinjas {
		go attack(ninjaBhau)
	}

	time.Sleep(time.Second)

	fmt.Println("ending go routine demo....")

	fmt.Println()
	fmt.Println()

	smokeSignal := make(chan bool)

	startChannel := time.Now()

	fmt.Println("Starting channels demo......", startChannel)

	fmt.Println()

	defer func() {
		fmt.Println(time.Since(startChannel))
	}()

	evilNin := "Hurry"
	go AttackNinja(evilNin, smokeSignal)

	fmt.Println(<-smokeSignal)

	fmt.Println("ending channels demo......")

}

func attack(ninjaBhau string) {
	fmt.Println("The ninja attack by", ninjaBhau)

}

func AttackNinja(target string, isAttacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Attack Ninja by", target)
	isAttacked <- true
}
