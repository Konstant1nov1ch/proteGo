package base

import (
	"github.com/ichiban/prolog"
	"log"
	"os"
)

func MakeNewBase() (*prolog.Interpreter, error) {
	p := prolog.New(nil, os.Stdout)

	if err := p.Exec(`
mob(creeper).
mob(zombie).
mob(ghast).
mob(skeleton).
mob(piglin).
mob(dragon).
mob(spider).
mob(steve).
mob(pig).
mob(cow).
mob(horse).
mob(cammel).

damage(creeper, 30).
damage(zombie, 10).
damage(ghast, 20).
damage(skeleton, 15).
damage(piglin, 25).
damage(dragon, 50).
damage(spider, 5).
damage(steve, 5).
damage(pig, 0).
damage(cow, 2).
damage(horse, 5).
damage(cammel, 5).

biom(plains, [pig, cow, horse, camel]).
biom(forest, [pig, cow, horse, camel, creeper, skeleton]).
biom(desert, [camel, skeleton]).
biom(jungle, [parrot, ocelot, creeper, skeleton]).
biom(swamp, [slime, witch, creeper, spider]).
biom(ocean, [dolphin, squid]).
biom(savanna, [zebra, giraffe, lion, camel]).
biom(taiga, [polar_bear, wolf, fox, creeper, skeleton]).
biom(mushroom_forest, [mooshroom, witch, creeper]).
biom(nether, [blaze, ghast, wither_skeleton]).
biom(end, [enderman, ender_dragon]).

friend(cow, horse).
friend(horse, cow).
friend(pig, piglin).
friend(piglin, pig).
friend(parrot, ocelot).
friend(ocelot, parrot).
friend(zebra, giraffe).
friend(giraffe, zebra).
friend(polar_bear, wolf).
friend(wolf, polar_bear).
friend(fox, wolf).
friend(wolf, fox).
friend(mooshroom, mushroom_cow).
friend(mushroom_cow, mooshroom).
friend(dolphin, squid).
friend(squid, dolphin).

enemy(creeper, player).
enemy(zombie, player).
enemy(ghast, player).
enemy(skeleton, player).
enemy(piglin, player).
enemy(dragon, player).
enemy(spider, player).
enemy(steve, player).
enemy(enderman, player).

who_is_stronger(X, Y, Stronger) :- damage(X, XDamage), damage(Y, YDamage),
    (XDamage > YDamage -> Stronger = X ; Stronger = Y).

can_attack(X, 'да') :- mob(X), damage(X, Damage), Damage > 0.
can_attack(_, 'нет').

can_explode(X, да) :- mob(X), X == creeper, !.
can_explode(_, нет).

can_fly(X, да) :- mob(X), (X == ghast ; X == dragon), !.
can_fly(_, нет).

can_teleport(X, да) :- mob(X), X == enderman, !.
can_teleport(_, нет).


is_friendly(X, Y, 'да') :- friend(X, Y).
is_friendly(Y, X, 'да') :- friend(X, Y).
is_friendly(_, _, 'нет').

is_enemy(X, Y, 'да') :- enemy(X, Y).
is_enemy(Y, X, 'да') :- enemy(X, Y).
is_enemy(_, _, 'нет').
`); err != nil {
		log.Fatalf("%v: \n", err)
		return nil, err
	}
	return p, nil
}
