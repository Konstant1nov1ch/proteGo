package main

import (
	"awesomeTestSolid/internal/base"
	"awesomeTestSolid/internal/model"
	"bufio"
	"fmt"
	"github.com/ichiban/prolog"
	"os"
	"strings"
)

func main() {
	p, err := base.MakeNewBase()
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	structures := map[string]interface{}{
		"1": &model.Mobs,
		"2": &model.Mob,
		"3": &model.Damage,
		"4": &model.BiomeMobs,
		"5": &model.IsFriendly,
		"6": &model.IsEnemy,
		"7": &model.AttackableMobs,
		"8": &model.StrengthComparison,
		"9": &model.Can,
	}

	for {
		fmt.Println("Select a query (1-8) or type 'exit' to quit:")
		fmt.Println("1. damage(zombie, X).")
		fmt.Println("2. biom(forest, Mobs).")
		fmt.Println("3. is_friendly(cow, creeper, X).")
		fmt.Println("4. is_enemy(creeper, player, X).")
		fmt.Println("5. can_attack(pig, X).")
		fmt.Println("6. can_fly(dragon, X).")
		fmt.Println("7. can_teleport(steve, X).")
		fmt.Println("8. who_is_stronger(dragon, skeleton, X).")

		fmt.Print("Enter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		switch input {
		case "1":
			runQuery(structures["3"], p, `damage(zombie, X).`)
		case "2":
			runQuery(structures["1"], p, `biom(forest, Mobs).`)
		case "3":
			runQuery(structures["9"], p, `is_friendly(cow, creeper, X).`)
		case "4":
			runQuery(structures["9"], p, `is_enemy(creeper, player, X).`)
		case "5":
			runQuery(structures["9"], p, `can_attack(cow, X).`)
		case "6":
			runQuery(structures["9"], p, `can_fly(dragon, X).`)
		case "7":
			runQuery(structures["9"], p, `can_teleport(pig, X).`)
		case "8":
			runQuery(structures["8"], p, `who_is_stronger(dragon, skeleton, X).`)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 8 or 'exit'.")
		}
	}
}

func runQuery(result interface{}, p *prolog.Interpreter, query string) {
	sols, err := p.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	for sols.Next() {
		if err := sols.Scan(result); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Result: %v\n", result)
	}

	if err := sols.Err(); err != nil {
		fmt.Println(err)
	}
}
