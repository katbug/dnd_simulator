package creature

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Attack struct {
	Weapon    string `json:"weapon"`
	ToHit     int    `json:"toHit"`
	DamageDie int    `json:"damageDie"`
	NumDice   int    `json:"numDice"`
	DamageMod int    `json:"damageMod"`
}

type StatBlock struct {
	Str int `json:"str"`
	Dex int `json:"dex"`
	Con int `json:"con"`
	Wis int `json:"wis"`
	Int int `json:"int"`
	Cha int `json:"cha"`
}

type Creature interface {
	GetName() string
	GetHP() int
	GetAC() int
	GetStats() StatBlock
	GetAttacks() []Attack
	GetType() bool
	PrintAttacks()
	PrintAttack(attackIdx int)
	Attack(idx int, target CreatureImpl)
	TakeDamage(damage int) bool
}

type CreatureImpl struct {
	Name    string    `json:"name"`
	HP      int       `json:"totalHP"`
	AC      int       `json:"ac"`
	Stats   StatBlock `json:"stats"`
	Attacks []Attack  `json:"attacks"`
	IsEnemy bool      `json:"isEnemy"`
}

func (c *CreatureImpl) GetName() string {
	return c.Name
}

func (c *CreatureImpl) GetHP() int {
	return c.HP
}
func (c *CreatureImpl) GetAC() int {
	return c.AC
}
func (c *CreatureImpl) GetStats() StatBlock {
	return c.Stats
}
func (c *CreatureImpl) GetAttacks() []Attack {
	return c.Attacks
}

func (c *CreatureImpl) GetType() bool {
	return c.IsEnemy
}

func (c *CreatureImpl) PrintAttacks() {
	fmt.Println(c.Name, " has the following attacks available: ")
	for i, a := range c.Attacks {
		fmt.Println(strconv.Itoa(i)+". ", a.Weapon)
	}
}

func (c *CreatureImpl) PrintAttack(attackIdx int) {
	a := c.Attacks[attackIdx]
	fmt.Println("Weapon:", a.Weapon, "+"+strconv.Itoa(a.ToHit))
	fmt.Println("Damage:", strconv.Itoa(a.NumDice)+"d"+strconv.Itoa(a.DamageDie), "+"+strconv.Itoa(a.DamageMod))
}

func (c *CreatureImpl) Attack(idx int, target *CreatureImpl) {
	attack := c.Attacks[idx]
	rand.Seed(time.Now().UnixNano())

	toHit := rand.Intn(20) + 1 + attack.ToHit
	if toHit < target.GetAC() {
		fmt.Println("Miss!")
		return
	}

	damage := attack.DamageMod
	for i := 0; i < attack.NumDice; i++ {
		damage += rand.Intn(attack.DamageDie) + 1
	}
	fmt.Println("Damage:", damage)
	dead := target.TakeDamage(damage)
	fmt.Println("Target's remaining HP:", target.GetHP())
	if dead {
		fmt.Println("Target has fallen")
	}

}

func (c *CreatureImpl) TakeDamage(damage int) bool {
	c.HP = c.HP - damage
	if c.HP <= 0 {
		c.HP = 0
		return true
	}
	return false
}
