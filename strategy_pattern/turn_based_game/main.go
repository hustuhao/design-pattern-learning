package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 本例实现了一个简单的回合制游戏

// Character 属性
type Character struct {
	Name            string
	Attributes      map[string]int
	AttackStrategy  AttackStrategy
	DefenseStrategy DefenseStrategy
	MagicStrategy   MagicStrategy
}

// AttackStrategy 处理攻击的策略
type AttackStrategy interface {
	ExecuteAttack(attacker *Character, defender *Character) (int, bool)
}

// DefenseStrategy 处理防御的策略
type DefenseStrategy interface {
	ExecuteDefense(attacker *Character, defender *Character) int
}

// MagicStrategy 处理魔法攻击的策略
type MagicStrategy interface {
	ExecuteMagic(attacker *Character, defender *Character) (int, bool)
}

// PhysicalAttackStrategy 物理攻击策略
type PhysicalAttackStrategy struct{}

func (a *PhysicalAttackStrategy) ExecuteAttack(attacker *Character, defender *Character) (int, bool) {
	damage := attacker.Attributes["attack_points"] - defender.Attributes["defense_points"]
	if rand.Intn(100) < attacker.Attributes["luck_points"] {
		damage *= 2
		return damage, true
	}
	return damage, false
}

// CounterAttackStrategy 反击策略
type CounterAttackStrategy struct{}

func (d *CounterAttackStrategy) ExecuteDefense(attacker *Character, defender *Character) int {
	damage := attacker.Attributes["attack_points"] - defender.Attributes["defense_points"]
	if rand.Intn(100) < defender.Attributes["luck_points"] {
		damage = int(float32(damage) * 0.5)
	}
	return damage
}

// MagicAttackStrategy 魔法攻击策略
type MagicAttackStrategy struct{}

func (m *MagicAttackStrategy) ExecuteMagic(attacker *Character, defender *Character) (int, bool) {
	damage := attacker.Attributes["magic_points"] - defender.Attributes["magic_defense_points"]
	if rand.Intn(100) < attacker.Attributes["luck_points"] {
		damage *= 2
		return damage, true
	}
	return damage, false
}

// Character 构造函数
func NewCharacter(name string, attributes map[string]int, attackStrategy AttackStrategy, defenseStrategy DefenseStrategy, magicStrategy MagicStrategy) *Character {
	return &Character{
		Name:            name,
		Attributes:      attributes,
		AttackStrategy:  attackStrategy,
		DefenseStrategy: defenseStrategy,
		MagicStrategy:   magicStrategy,
	}
}

// Attack 让 attacker 攻击 defender
func Attack(attacker *Character, defender *Character) (int, bool) {
	return attacker.AttackStrategy.ExecuteAttack(attacker, defender)
}

// MagicAttack 让 attacker 使用魔法攻击 defender
func MagicAttack(attacker *Character, defender *Character) (int, bool) {
	return attacker.MagicStrategy.ExecuteMagic(attacker, defender)
}

// AttackLog 每次攻击的日志
type AttackLog struct {
	AttackerName    string `json:"attacker"`
	DefenderName    string `json:"defender"`
	DamageValue     int    `json:"damage"`
	Critical        bool   `json:"critical"`
	LifestealAmount int    `json:"lifesteal"`
}

// GameLog 所有攻击的日志记录
type GameLog struct {
	Attacks []AttackLog `json:"attacks"`
}

func (g *GameLog) AddAttack(attacker *Character, defender *Character, damageValue int, critical bool, lifestealAmount int) {
	attack := AttackLog{
		AttackerName:    attacker.Name,
		DefenderName:    defender.Name,
		DamageValue:     damageValue,
		Critical:        critical,
		LifestealAmount: lifestealAmount,
	}

	g.Attacks = append(g.Attacks, attack)
}

// PlayGame 游戏主逻辑
func PlayGame(player1 *Character, player2 *Character) GameLog {
	gameLog := GameLog{}
	attacker := player1
	defender := player2
	for {
		// 随机选择攻击方式
		var damage int
		var critical bool
		switch rand.Intn(4) {
		case 1:
			damage, critical = MagicAttack(attacker, defender)
		default:
			damage, critical = Attack(attacker, defender)

		}
		lifeStealAmount := 0
		if critical {
			lifeStealAmount = int(float32(attacker.Attributes["attack_points"]) * 0.1)
			attacker.Attributes["health_points"] += lifeStealAmount
		}
		defender.Attributes["health_points"] -= damage
		gameLog.AddAttack(attacker, defender, damage, critical, lifeStealAmount)
		if defender.Attributes["health_points"] <= 0 {
			break
		}
		attacker, defender = defender, attacker
	}
	return gameLog
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 玩家1
	player1Attributes := map[string]int{ // 初始化人数属性
		"health_points":        100,
		"attack_points":        20,
		"defense_points":       15,
		"magic_points":         15,
		"magic_defense_points": 10,
		"agility_points":       10,
		"luck_points":          5,
	}
	player1AttackStrategy := &PhysicalAttackStrategy{}
	player1DefenseStrategy := &CounterAttackStrategy{}
	player1MagicStrategy := &MagicAttackStrategy{}
	player1 := NewCharacter("player1", player1Attributes, player1AttackStrategy, player1DefenseStrategy, player1MagicStrategy)
	// 玩家2
	player2Attributes := map[string]int{
		"health_points":        120,
		"attack_points":        15,
		"defense_points":       10,
		"magic_points":         20,
		"magic_defense_points": 15,
		"agility_points":       10,
		"luck_points":          5,
	}
	player2AttackStrategy := &PhysicalAttackStrategy{}
	player2MagicStrategy := &MagicAttackStrategy{}
	player2DefenseStrategy := &CounterAttackStrategy{}
	player2 := NewCharacter("player2", player2Attributes, player2AttackStrategy, player2DefenseStrategy, player2MagicStrategy)

	// 开始游戏
	gameLog := PlayGame(player1, player2)
	fmt.Println(gameLog)
}
