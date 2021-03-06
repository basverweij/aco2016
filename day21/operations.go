package main

import "fmt"

type operation interface {
	perform(pwd *password)
	unperform(pwd *password)
}

type swapPosition struct {
	from, to int
}

func (sp *swapPosition) String() string {
	return fmt.Sprintf(
		"swap position %d with position %d",
		sp.from,
		sp.to)
}

func (sp *swapPosition) perform(pwd *password) {
	pwd.swap(sp.from, sp.to)
}

func (sp *swapPosition) unperform(pwd *password) {
	pwd.swap(sp.to, sp.from)
}

type swapLetter struct {
	fromLetter, toLetter byte
}

func (sl *swapLetter) String() string {
	return fmt.Sprintf(
		"swap letter %s with letter %s",
		string(sl.fromLetter),
		string(sl.toLetter))
}

func (sl *swapLetter) perform(pwd *password) {
	var from, to int

	for i, l := range pwd.value {
		if l == sl.fromLetter {
			from = i
		} else if l == sl.toLetter {
			to = i
		}
	}

	pwd.swap(from, to)
}

func (sl *swapLetter) unperform(pwd *password) {
	var from, to int

	for i, l := range pwd.value {
		if l == sl.fromLetter {
			from = i
		} else if l == sl.toLetter {
			to = i
		}
	}

	pwd.swap(to, from)
}

type rotate struct {
	rotateLeft bool
	steps      int
}

func (r *rotate) String() string {
	var dir string
	if r.rotateLeft {
		dir = "left"
	} else {
		dir = "right"
	}

	suffix := ""
	if r.steps != 1 {
		suffix = "s"
	}

	return fmt.Sprintf(
		"rotate %s %d step%s",
		dir,
		r.steps,
		suffix)
}

func (r *rotate) perform(pwd *password) {
	pwd.rotate(r.rotateLeft, r.steps)
}

func (r *rotate) unperform(pwd *password) {
	pwd.rotate(!r.rotateLeft, r.steps)
}

type rotateLetter struct {
	letter byte
}

func (rl *rotateLetter) String() string {
	return fmt.Sprintf(
		"rotate based on position of letter %s",
		string(rl.letter))
}

func (rl *rotateLetter) perform(pwd *password) {
	var steps int

	for i, l := range pwd.value {
		if l == rl.letter {
			steps = i + 1
			break
		}
	}

	if steps >= 5 {
		steps++
	}

	pwd.rotate(false, steps)
}

var unperformRotateLetter = map[int]int{
	1: 1,
	3: 2,
	5: 3,
	7: 4,
	2: 6,
	4: 7,
	6: 0,
	0: 1,
}

func (rl *rotateLetter) unperform(pwd *password) {
	var steps int

	for i, l := range pwd.value {
		if l == rl.letter {
			steps = i
			break
		}
	}

	pwd.rotate(true, unperformRotateLetter[steps])
}

type reverse struct {
	from, to int
}

func (r *reverse) String() string {
	return fmt.Sprintf(
		"reverse positions %d through %d",
		r.from,
		r.to)
}

func (r *reverse) perform(pwd *password) {
	pwd.reverse(r.from, r.to)
}

func (r *reverse) unperform(pwd *password) {
	pwd.reverse(r.to, r.from)
}

type move struct {
	from, to int
}

func (m *move) String() string {
	return fmt.Sprintf(
		"move position %d to position %d",
		m.from,
		m.to)
}

func (m *move) perform(pwd *password) {
	pwd.move(m.from, m.to)
}

func (m *move) unperform(pwd *password) {
	pwd.move(m.to, m.from)
}
