package games

type Fourwins struct {
	id int
}

func (game *Fourwins) GetConfig() GameConfig {
	return GameConfig{4}
}

func (game *Fourwins) Move() {

}
