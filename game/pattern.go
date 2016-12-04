package game

type PatternCollection struct {
	Patterns map[int][]*Pattern
}

type Pattern struct {
	spawns       []*Spawn
	currentSpawn int
}

func NewPattern(spawns []*Spawn) *Pattern {
	var pattern = &Pattern{
		spawns:       spawns,
		currentSpawn: 0,
	}
	return pattern
}

func (p *Pattern) GetCurrentSpawn() *Spawn {
	return p.spawns[p.currentSpawn]
}

func (p *Pattern) OnLastSpawn() bool {
	return p.currentSpawn == len(p.spawns)-1
}

func (p *Pattern) ResetPattern() {
	p.currentSpawn = 0

}

func (p *Pattern) AdvancePattern() {
	p.currentSpawn++
}

type Spawn struct {
	ShapeType int
	Track     int
	Speed     int
	// how long before this spawn should be added to the collection (milliseconds)
	SpawnDelayMillis int
}

func NewSpawn(shapeType int, track int, speed int, spawnDelayMillis int) *Spawn {
	var spawn = &Spawn{
		ShapeType:        shapeType,
		Track:            track,
		Speed:            speed,
		SpawnDelayMillis: spawnDelayMillis,
	}
	return spawn
}

func NewSpawnDefaultSpeed(shapeType int, track int, spawnDelayMillis int) *Spawn {
	var spawn = &Spawn{
		ShapeType:        shapeType,
		Track:            track,
		Speed:            DefaultSpeed,
		SpawnDelayMillis: spawnDelayMillis,
	}
	return spawn
}
