package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/dnd_simulator/cmd"
	creature "github.com/dnd_simulator/creature"
)

type Attack struct {
	Weapon    string `json:"weapon"`
	ToHit     int    `json:"toHit"`
	DamageDie int    `json:"damageDie"`
	DamageMod int    `json:"damageMod"`
}

type Player interface {
	creature.Creature
}

type PlayerImpl struct {
	creature.CreatureImpl `json:"creature"`
}

type Enemy struct {
	creature.CreatureImpl `json:"creature"`
}

func loadCreatures(fileName string, destination *[]creature.CreatureImpl) {
	// Open our players jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened ", fileName)
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, destination)
	fmt.Println("players: ", destination)
}

type PosInit struct {
	Pos        int
	Initiative int
	Dex        int
}

func rollInitiative(players []creature.CreatureImpl, enemies []creature.CreatureImpl) *[]creature.CreatureImpl {
	total := append(players, enemies...)
	totalLen := len(total)
	order := make([]creature.CreatureImpl, totalLen)

	rand.Seed(time.Now().UnixNano())

	sorting := make([]PosInit, totalLen)

	for i, val := range total {
		dex := val.GetStats().Dex
		val := rand.Intn(20) + 1 + dex
		sorting[i] = PosInit{i, val, dex}
	}

	sort.Slice(sorting, func(i, j int) bool {
		if sorting[i].Initiative > sorting[j].Initiative {
			return true
		} else if sorting[i].Initiative < sorting[j].Initiative {
			return false
		}
		return sorting[i].Dex > sorting[j].Dex
	})

	for i, v := range sorting {
		order[i] = total[v.Pos]
	}

	return &order
}

func getChoice(scanner *bufio.Scanner) int {
	goodInput := false
	choice := -1
	var err error

	for !goodInput {
		scanner.Scan()
		choice, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Incorrect input. Please choose a number.")
		} else {
			goodInput = true
		}
	}
	return choice
}

func printTargets(targets []creature.CreatureImpl) {
	for i, c := range targets {
		fmt.Println(strconv.Itoa(i) + ". " + c.GetName())
	}
}

func attack(creature, target creature.Creature, attackIdx int) {

}

func main() {

	cmd.Execute()

	var players []creature.CreatureImpl
	loadCreatures("pcs.json", &players)

	var enemies []creature.CreatureImpl
	loadCreatures("enemies.json", &enemies)

	fmt.Println("players: ", players)
	fmt.Println(players[0].GetAttacks())

	e := Enemy{creature.CreatureImpl{"bob", 10, 10, creature.StatBlock{}, []creature.Attack{}, true}}
	fmt.Println("enemy: ", e.GetName())

	order := rollInitiative(players, enemies)
	fmt.Println("order: ", order)

	total := len(*order)
	currIdx := 0

	scanner := bufio.NewScanner(os.Stdin)

	for {
		curr := (*order)[currIdx]
		fmt.Println("Creature: ", curr.GetName(), " is up.")
		curr.PrintAttacks()
		fmt.Println("Please select the desired attack by index.")
		attack := getChoice(scanner)
		fmt.Println("You have chosen attack: ")
		curr.PrintAttack(attack)

		fmt.Println("Please choose a target: ")
		//var target creature.CreatureImpl
		printTargets(*order)
		targetIdx := getChoice(scanner)
		target := &(*order)[targetIdx]
		/*if curr.GetType() == true {

			targetIdx := getChoice(scanner)
			target = players[targetIdx]
		} else {
			printTargets(enemies)
			targetIdx := getChoice(scanner)
			target = enemies[targetIdx]
		}*/

		fmt.Println("You have chosen target: ")
		fmt.Println(target.GetName())
		fmt.Println("Attacking...")
		curr.Attack(attack, target)
		fmt.Println("target: ", target)

		currIdx += 1
		currIdx = currIdx % total
	}

}
