package main

func RunLogic(spheres []Sphere) State {
	deads := 0
	for i := range spheres {
		if spheres[i].Hp <= 0 {
			spheres[i].Alive = false
			deads += 1
		}
	}

	switch deads {
	case len(spheres) - 1:
		return WON
	case len(spheres):
		return DRAW
	default:
		return FIGHTING
	}
}
