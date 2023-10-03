package model

var Mobs struct {
	Who []string `prolog:"Mobs"`
}

var Mob struct {
	Who string `prolog:"Mob"`
}
var Damage struct {
	Amount int `prolog:"X"`
}
var BiomeMobs struct {
	Mobs []string `prolog:"Mobs"`
}
var IsFriendly struct {
	IsFriendly bool `prolog:"true"`
}
var IsEnemy struct {
	IsEnemy bool `prolog:"true"`
}
var AttackableMobs struct {
	Who string `prolog:"X"`
}
var StrengthComparison struct {
	Stronger string `prolog:"X"`
}
var Can struct {
	X string `prolog:"X"`
}
