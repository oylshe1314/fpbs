package util

func NumbersConvert1[T1 Number, T2 Number](s1 []T1, s2 *[]T2) {
	*s2 = make([]T2, len(s1))
	for i := range s1 {
		(*s2)[i] = T2(s1[i])
	}
}

func NumbersConvert2[T1 Number, T2 Number](s1 []T1, t2 T2) []T2 {
	var s2 []T2
	NumbersConvert1(s1, &s2)
	return s2
}
