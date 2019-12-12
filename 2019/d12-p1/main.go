package main

import (
	"fmt"
)

func main() {

	//Puzzle Input
	moonLocations := []xyz{xyz{13, 9, 5}, xyz{8, 14, -2}, xyz{-5, 4, 11}, xyz{2, -6, 1}}

	//Print the result
	fmt.Println(calculateTotalEnergy(moonLocations, 1000))
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func step(moons []moon) ([]moon, int) {
	energy := 0

	calced := make(map[string]bool)

	//Calculate Gravity
	for i := 0; i < len(moons); i++ {
		for j := 0; j < len(moons); j++ {
			if i == j {
				//Pretty sure we could get away with calculating gravity with it's self as it shouldn't affect the velocity
				continue
			}
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

		energy += calculateEnergy(moons[i])

	}
	// fmt.Println("Velocity Calc:", moons)

	return moons, energy
}

func calculateTotalEnergy(moonLocations []xyz, steps int) int {
	totalEnergy := 0
	energy := 0

	fmt.Println("\n\ncalculateTotalEnergy", steps)

	moons := []moon{}
	for _, m := range moonLocations {
		//Moons always start with velcoity 0
		moons = append(moons, moon{m, xyz{0, 0, 0}})

	}
	fmt.Println(0, moons)
	//1000 steps
	for i := 0; i < steps; i++ {
		moons, energy = step(moons)
		// fmt.Println("Step:", i)
		// for j, m := range moons {
		// 	fmt.Println(j, m)
		// }
		// fmt.Println(i+1, moons, energy)
		totalEnergy = energy //don't want to sum up steps to just update this every time
	}

	return totalEnergy
}
