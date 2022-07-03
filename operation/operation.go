package operation

import (
	"errors"
	"math"
	"strings"

	"apirest.ofq/structs"
)

var kenobi = structs.CoordPos{X: -500, Y: -200}
var skywalker = structs.CoordPos{X: 100, Y: -100}
var sato = structs.CoordPos{X: 500, Y: 100}

func GetLocation(distances ...float32) (x float32, y float32, err error) {
	if len(distances) < 3 {
		err = errors.New("hay menos de 3 distancias")
		return
	}
	//Distance equation kenobi - skywalker
	a, b, c := calcDistance2Satellites(kenobi, skywalker, distances[0], distances[1])
	//Distance equation skywalker - sato
	e, f, g := calcDistance2Satellites(skywalker, sato, distances[1], distances[2])
	//Finding Cramer's rule determinants
	d := a*f - b*e
	if d == 0.0 {
		err = errors.New("distancia menor a cero")
		return
	}

	x = (c*f - b*g) / d
	y = (a*g - c*e) / d
	return
}

func calcDistance2Satellites(coordSat1 structs.CoordPos, coordSat2 structs.CoordPos, distanceSat1 float32, distanceSat2 float32) (a, b, c float32) {
	a = -2*coordSat1.X + 2*coordSat2.X
	b = -2*coordSat1.Y + 2*coordSat2.Y
	c = float32(math.Pow(float64(distanceSat1), 2) - math.Pow(float64(distanceSat2), 2) - math.Pow(float64(coordSat1.X), 2) + math.Pow(float64(coordSat2.X), 2) - math.Pow(float64(coordSat1.Y), 2) + math.Pow(float64(coordSat2.Y), 2))
	return
}

func GetMessage(messages ...[]string) (msg string, err error) {
	mergedMessage := make([]string, 0)
	messagePart := ""
	for i, v0 := range messages {
		actualMessage := v0
		for j, v1 := range actualMessage {
			messagePart = v1
			if v1 == "" {
				for k := (i + 1); k < len(messages); k++ {
					messagePart = messages[k][j]
					if messagePart != "" {
						break
					}
				}
			}
			mergedMessage = append(mergedMessage, messagePart)
		}
		if len(mergedMessage) == len(actualMessage) {
			break
		}
	}
	resp := strings.Join(mergedMessage, " ")
	if strings.Trim(resp, " ") == "" {
		err = errors.New("mensaje no desifrado")
	}
	return resp, err
}

func GetCoordSat(name string) structs.CoordPos {
	var coord structs.CoordPos
	switch name {
	case "kenobi":
		return kenobi
	case "skywalker":
		return skywalker
	case "sato":
		return sato
	default:
		return coord
	}
}
