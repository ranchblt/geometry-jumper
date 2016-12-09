package game

import "encoding/json"

type PatternCollection struct {
	Patterns map[string][]*Pattern
}

func PatternCollectionFromJSON(data []byte) *PatternCollection {
	var patternCollection PatternCollection
	error := json.Unmarshal(data, &patternCollection)
	if error != nil {
		panic(error)
	}

	return &patternCollection
}

type Pattern struct {
	SpawnGroups []*SpawnGroup
}

func NewPattern(spawnGroups []*SpawnGroup) *Pattern {
	var pattern = &Pattern{
		SpawnGroups: spawnGroups,
	}
	return pattern
}

type SpawnGroup struct {
	Spawns []*Spawn
	// how long since the start of the pattern to spawn this group of shapes
	SpawnTimeMillis int
}

func NewSpawnGroup(spawns []*Spawn, spawnTimeMillis int) *SpawnGroup {
	var spawnGroup = &SpawnGroup{
		Spawns:          spawns,
		SpawnTimeMillis: spawnTimeMillis,
	}
	return spawnGroup
}

type Spawn struct {
	ShapeType int
	Track     int
	Speed     int
}

func NewSpawn(shapeType int, track int, speed int) *Spawn {
	var spawn = &Spawn{
		ShapeType: shapeType,
		Track:     track,
		Speed:     speed,
	}
	return spawn
}

func NewSpawnDefaultSpeed(shapeType int, track int) *Spawn {
	var spawn = &Spawn{
		ShapeType: shapeType,
		Track:     track,
		Speed:     DefaultSpeed,
	}
	return spawn
}
