module goWordle/game

go 1.21.5

replace goWordle/wordChoice => ../wordChoice

require (
	goWordle/checkGuess v0.0.0-00010101000000-000000000000
	goWordle/checkWord v0.0.0-00010101000000-000000000000
	goWordle/wordChoice v0.0.0-00010101000000-000000000000
)

replace goWordle/checkGuess => ../checkGuess

replace goWordle/checkWord => ../checkWord
