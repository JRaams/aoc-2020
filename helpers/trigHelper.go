package helpers

import "math"

// GetManhattanDistInt returns the Manhattan distance between a point defined with int values (x,y) and (0,0)
func GetManhattanDistInt(x int, y int) int {
	return int(GetManhattanDistFloat64(float64(x), float64(y)))
}

// GetManhattanDistFloat64 returns the Manhattan distance between a point defined with float64 values (x,y) and (0,0)
func GetManhattanDistFloat64(x float64, y float64) float64 {
	return math.Abs(x) + math.Abs(y)
}

// RotateCartesianIntCoordsByDegrees rotates a set of integer coordinates 'x' and 'y' by 'degrees'
func RotateCartesianIntCoordsByDegrees(x int, y int, degrees int) (int, int) {
	newX, newY := RotateCartesianFloat64CoordsByDegrees(float64(x), float64(y), float64(degrees))
	return int(math.Round(newX)), int(math.Round(newY))
}

// RotateCartesianFloat64CoordsByDegrees rotates a set of float64 coordinates 'x' and 'y' by 'degrees'
func RotateCartesianFloat64CoordsByDegrees(x float64, y float64, degrees float64) (float64, float64) {
	radians := degrees * (math.Pi / 180)
	newX := x*math.Cos(radians) + y*math.Sin(radians)
	newY := y*math.Cos(radians) - x*math.Sin(radians)
	return newX, newY
}
