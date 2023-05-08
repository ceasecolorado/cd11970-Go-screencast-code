package main

type Baseball struct {
	mass         int
	accelatation int
}

type Football struct {
}

func (b Baseball) getForce() int {
	return b.mass * b.accelatation
}

func (b Football) getForce() int {
	return 50
}

type Force interface {
	getForce() int
}

func compareForce(a, b Force) bool {
	return a.getForce() > b.getForce()
}

func main() {
	b := Baseball{2, 20}
	f := Football{}
	println(compareForce(b, f))
}
