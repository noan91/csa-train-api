package traindiary

type ExcerciseTarget int

const (
	Legs ExcerciseTarget = iota
	Back
	Chest
	Shoulders
	Biceps
	Triceps
	Core
)

type Excercise struct {
	Name     string
	Target   ExcerciseTarget
	WeightOn bool
}

type TrainSetCounter int

const (
	Reps TrainSetCounter = iota
	Time
)

type TrainSet struct {
	Excercise Excercise
	Counter   TrainSetCounter
	Value     float64
	Intencity float64
}
