// This package contains info about player and it's functions
package player

type Player struct {
	name           string
	no_of_attempts uint16
}

func CreatePlayer(nm string) *Player {

	return &Player{
		name:           nm,
		no_of_attempts: 0,
	}
}

func (p *Player) GetNoOfAttempts() uint16 {
	return p.no_of_attempts
}

func (p *Player) IncreaseNoOfAttempts() {
	p.no_of_attempts++
}
