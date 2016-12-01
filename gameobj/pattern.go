package gameobj

type PatternCollection struct {
	Patterns map[int][]*Pattern
}

type Pattern struct {
	spawns []*Spawn
}

func NewPattern(spawns []*Spawn) *Pattern {
	var pattern = &Pattern{
		spawns: spawns,
	}
	return pattern
}

func (p *Pattern) GetCurrentSpawn(currentSpawn int) *Spawn {
	return p.spawns[currentSpawn]
}

func (p *Pattern) SpawnReady(currentSpawn int, timer int64) bool {
	return p.spawns[currentSpawn].NextSpawnDelay <= timer
}

func (p *Pattern) OnLastSpawn(currentSpawn int) bool {
	return currentSpawn == len(p.spawns)
}

type Spawn struct {
	ShapeType int
	Track     int
	Speed     int
	// how long before the next spawn should be added to the collection
	// TODO this is in seconds right now, but we should probably be able to spawn in split seconds of some measurement
	NextSpawnDelay int64
}

func NewSpawn(shapeType int, track int, speed int, nextSpawnDelay int64) *Spawn {
	var spawn = &Spawn{
		ShapeType:      shapeType,
		Track:          track,
		Speed:          speed,
		NextSpawnDelay: nextSpawnDelay,
	}
	return spawn
}

func NewEndcapSpawn(shapeType int, track int, speed int) *Spawn {
	var spawn = &Spawn{
		ShapeType:      shapeType,
		Track:          track,
		Speed:          speed,
		NextSpawnDelay: EndOfPatternSpawnDelay,
	}
	return spawn
}

func NewSpawnDefaultSpeed(shapeType int, track int, nextSpawnDelay int64) *Spawn {
	var spawn = &Spawn{
		ShapeType:      shapeType,
		Track:          track,
		Speed:          DefaultSpeed,
		NextSpawnDelay: nextSpawnDelay,
	}
	return spawn
}

func NewEndcapSpawnDefaultSpeed(shapeType int, track int) *Spawn {
	var spawn = &Spawn{
		ShapeType:      shapeType,
		Track:          track,
		Speed:          DefaultSpeed,
		NextSpawnDelay: EndOfPatternSpawnDelay,
	}
	return spawn
}
