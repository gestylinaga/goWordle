module goWordle

go 1.21.6

replace goWordle/checkGuess => ./checkGuess

replace goWordle/checkWord => ./checkWord

replace goWordle/wordChoice => ./wordChoice

require (
	goWordle/checkGuess v0.0.0-00010101000000-000000000000
	goWordle/checkWord v0.0.0-00010101000000-000000000000
	goWordle/wordChoice v0.0.0-00010101000000-000000000000
)
