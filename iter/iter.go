package iter

func Iter(c string, vr int) string {
	var rep string
	for i := 0; i < vr; i++ {
		rep += c
	}
	return rep
}