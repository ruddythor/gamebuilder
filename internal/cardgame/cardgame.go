package cardgame

type CardGame interface {
	Play()
	Initialize(Player)
	// ShowResults()
	// GetPlayerInput(*io.Reader)
	EvaluateWinRules() string
	ApplyRules()
	Deal()
	EndTurn()
	StartTurn()
}

type Player interface {
	GetName() string
	GetBankRoll() int
	// Hand
	// GetName() string
}

type Playerst struct {
	Name     string
	Hand     Hand
	BankRoll int
}

type Game struct {
	GameDeck Deck
	Player   Playerst
	Dealer   Playerst
	Rulebook *Rulebook
	// Board Board
}

type Rulebook struct {
	WinRules   []RuleCondition
	TurnRules  []RuleCondition
	SetupRules []RuleCondition
}

type Rule struct {
	RuleCondition
}

type RuleCondition interface {
	Condition(game interface{}) bool
	Name() string
}
