package move

var moves = map[string]int{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var goals = map[string]int{
	"X": Loss,
	"Y": Draw,
	"Z": Win,
}

var wins = map[int]int{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var loses = map[int]int{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

const (
	Rock     int = 1
	Paper    int = 2
	Scissors int = 3
)

const (
	Win  int = 6
	Draw int = 3
	Loss int = 0
)

type Action interface {
	CalculateScore() int
}

type Move struct {
	Opponent int
	Self     int
}

type Outcome struct {
	Opponent int
	Goal     int
}

func NewMove(inputs []string) Action {
	return Move{Opponent: moves[inputs[0]], Self: moves[inputs[1]]}
}

func NewOutcome(inputs []string) Action {
	return Outcome{Opponent: moves[inputs[0]], Goal: goals[inputs[1]]}
}

func (m Move) CalculateScore() int {
	score := m.Self
	if m.Self == m.Opponent {
		score += Draw
	} else if wins[m.Self] == m.Opponent {
		score += Win
	} else {
		score += Loss
	}
	return score
}

func (o Outcome) CalculateScore() int {
	score := o.Goal
	if o.Goal == Draw {
		score += o.Opponent
	} else if o.Goal == Win {
		score += loses[o.Opponent]
	} else {
		score += wins[o.Opponent]
	}
	return score
}
