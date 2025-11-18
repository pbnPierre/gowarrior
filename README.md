# gowarrior

This is a game designed to teach the Go language in a fun and interactive way.

You play as a warrior climbing a tall tower to reach the precious Treasure at the top level. On each floor you need to write a Go script to instruct the warrior to battle enemies, rescue captives, and reach the stairs.

You have some idea of what each floor contains, but you never know for certain what will happen. You must give the Warrior enough intelligence up-front to find his own way and defeat foes.

## Getting Started

TODO

## Scoring

Your objective is to not only reach the stairs, but to get the highest score you can. There are many ways you can earn points on a level.

* defeat an enemy to add his max health to your score
* rescue a captive to earn 20 points
* pass the level within the bonus time to earn the amount of bonus time remaining
* defeat all enemies and rescue all captives to receive a 20% overall bonus

A total score is kept as you progress through the levels. When you pass a level, that score is added to your total.

Don't be too concerned about scoring perfectly in the beginning. After you reach the top of the tower you will be able to re-run the tower and fine-tune your warrior to get the highest score. See the Epic Mode below for details.

## Perspective

TODO

## Commanding the Warrior

TODO

## Spaces

TODO

## Golem

TODO

## Epic Mode

Once you reach the top of the tower, you will enter epic mode. When running gowarrior again it will run your current player.rb through all levels in the tower without stopping.

Your warrior will most likely not succeed the first time around, so use the -l option on levels you are having difficulty or want to fine-tune the scoring.

```sh
TODO
```

Once your warrior reaches the top again you will receive an average grade, along with a grade for each level. The grades from best to worst are S, A, B, C, D and F. Try to get S on each level for the ultimate score.

Note: I'm in the process of fine-tuning the grading system. If you find the `S` grade to be too easy or too difficult to achieve on a given level, please add an issue for this on GitHub.

## Tips

TODO

## Other versions

This program is a direct inspiration of the brilliant [Ruby warrior](https://github.com/ryanb/ruby-warrior)

You can find other version of this project in a variety of languages

* [Godot](https://github.com/lkhedlund/godot-warrior)
* [PHP](https://github.com/yandod/php-warrior)
* [Python](https://github.com/arbylee/python-warrior)