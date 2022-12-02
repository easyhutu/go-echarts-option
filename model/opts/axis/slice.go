package axis

type Slices []Single

func (s Slices) Update(singles ...Single) {
	if len(singles) == 0 {
		panic(singles)
	}
	for idx, single := range singles {
		s[idx] = single
	}
}

func NewSlices(singles ...Single) Axis {
	slices := Slices{}
	slices = append(slices, singles...)
	return &slices
}
