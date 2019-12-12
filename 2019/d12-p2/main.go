package main

import (
	"fmt"
	"time"
)

func main() {

	//Puzzle Input
	moonLocations := []xyz{xyz{13, 9, 5}, xyz{8, 14, -2}, xyz{-5, 4, 11}, xyz{2, -6, 1}}

	//Print the result
	fmt.Println(calculateRepeatSteps(moonLocations))
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

type xyz struct {
	x int
	y int
	z int
}

func (p xyz) String() string {
	return fmt.Sprintf("<x=%d, y=%d,z=%d>", p.x, p.y, p.z)
}

type moon struct {
	pos xyz
	vel xyz
}

func (m moon) String() string {
	return fmt.Sprintf("pos=%v, vel=%v", m.pos, m.vel)
}
func calculateGravity(m1, m2 moon) (moon, moon) {

	if m1.pos.x == m2.pos.x {
		//No Change to x velocity
	} else if m1.pos.x > m2.pos.x {
		m1.vel.x--
		m2.vel.x++
	} else if m1.pos.x < m2.pos.x {
		m1.vel.x++
		m2.vel.x--
	}

	if m1.pos.y == m2.pos.y {
		//No Change to y velocity
	} else if m1.pos.y > m2.pos.y {
		m1.vel.y--
		m2.vel.y++
	} else if m1.pos.y < m2.pos.y {
		m1.vel.y++
		m2.vel.y--
	}

	if m1.pos.z == m2.pos.z {
		//No Change to z velocity
	} else if m1.pos.z > m2.pos.z {
		m1.vel.z--
		m2.vel.z++
	} else if m1.pos.z < m2.pos.z {
		m1.vel.z++
		m2.vel.z--
	}

	return m1, m2
}

func calculateEnergy(m moon) int {
	return (abs(m.pos.x) + abs(m.pos.y) + abs(m.pos.z)) * (abs(m.vel.x) + abs(m.vel.y) + abs(m.vel.z))
}

const MaxInt = int(^uint(0) >> 1)

func xKey(moons []moon) string {
	str := ""
	for _, m := range moons {
		str += fmt.Sprintf("(%d,%d)", m.pos.x, m.vel.x)
	}
	return str
}
func yKey(moons []moon) string {
	str := ""
	for _, m := range moons {
		str += fmt.Sprintf("(%d,%d)", m.pos.y, m.vel.y)
	}
	return str
}
func zKey(moons []moon) string {
	str := ""
	for _, m := range moons {
		str += fmt.Sprintf("(%d,%d)", m.pos.z, m.vel.z)
	}
	return str
}

func findPeriods(moons []moon) (int, int, int) {

	startTime := time.Now().Unix()

	stateX := make(map[string]bool)
	stateY := make(map[string]bool)
	stateZ := make(map[string]bool)

	periodX := 0
	periodY := 0
	periodZ := 0

	//State Keys
	sX := ""
	sY := ""
	sZ := ""

	for step := 0; step < MaxInt; step++ {
		sX = xKey(moons)
		// fmt.Println("xKey", sX)
		if periodX == 0 && stateX[sX] {
			periodX = step
			fmt.Println("periodX = ", periodX)
		} else {
			stateX[sX] = true
		}
		sY = yKey(moons)
		// fmt.Println("yKey", sY)
		if periodY == 0 && stateY[sY] {
			periodY = step
			fmt.Println("periodY = ", periodY)
		} else {
			stateY[sY] = true
		}
		sZ = zKey(moons)
		// fmt.Println("zKey", sZ)
		if periodZ == 0 && stateZ[sZ] {
			periodZ = step
			fmt.Println("periodZ = ", periodZ)
		} else {
			stateZ[sZ] = true
		}

		if periodX > 0 && periodY > 0 && periodZ > 0 {
			return periodX, periodY, periodZ
		}

		calced := make(map[string]bool)
		//Calculate Gravity
		for i := 0; i < len(moons); i++ {
			for j := 0; j < len(moons); j++ {
				if i < j && !calced[fmt.Sprintf("%d,%d", i, j)] {
					moons[i], moons[j] = calculateGravity(moons[i], moons[j])
					calced[fmt.Sprintf("%d,%d", i, j)] = true
					// fmt.Println("Gravity Calc:", i, j, moons)
				} else if j < i && !calced[fmt.Sprintf("%d,%d", j, i)] {
					//Probably should be able to eleminate this check but can't be bothered thinking it through...
					moons[i], moons[j] = calculateGravity(moons[i], moons[j])
					// fmt.Println("Gravity Calc B:", i, j, moons)
					calced[fmt.Sprintf("%d,%d", j, i)] = true
				}

			}

		}
		// fmt.Println("Gravity Calc:", moons)

		//Apply Velocity
		for i := 0; i < len(moons); i++ {
			moons[i].pos.x += moons[i].vel.x
			moons[i].pos.y += moons[i].vel.y
			moons[i].pos.z += moons[i].vel.z

		}
		// fmt.Println("Velocity Calc:", moons)

		if step > 0 && step%1000 == 0 {
			t := time.Now().Unix()
			fmt.Printf("Calculated %d %d/second\n", step, int64(step)/max((t-startTime), 1))
		}

	}

	return 0, 0, 0
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func calculateRepeatSteps(moonLocations []xyz) int {

	fmt.Println("\n\ncalculateRepeatSteps")

	moons := []moon{}
	for _, m := range moonLocations {
		//Moons always start with velcoity 0
		moons = append(moons, moon{m, xyz{0, 0, 0}})

	}
	fmt.Println(0, moons)

	pX, pY, pZ := findPeriods(moons)
	fmt.Println(pX, pY, pZ)

	//todo file Lowest Common Multiple
	result := LCM(pX, pY, pZ)

	return result
}
