package utils

func ValidateIIN(iin string) bool {
	if len(iin) < 12 {
		return false
	}

	// Разбираем цифры
	d := make([]int, 12)
	for i := 0; i < 12; i++ {
		c := iin[i]
		if c < '0' || c > '9' {
			return false
		}
		d[i] = int(c - '0')
	}

	s := 0
	for i := 0; i < 11; i++ {
		s += (i + 1) * d[i]
	}
	k := s % 11

	if k == 10 {
		s = 0
		for i := 0; i < 11; i++ {
			t := (i + 3) % 11
			if t == 0 {
				t = 11
			}
			s += t * d[i]
		}
		k = s % 11
		if k == 10 {
			return false
		}
		return k == d[11]
	}

	return k == d[11]
}
