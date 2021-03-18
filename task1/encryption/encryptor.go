package encryption

func Encrypt(text string, shift int32) (encrypted string){
	r := []rune(text)
	for i := 0; i < len(r); i++ {
		if r[i] >= 97 && r[i] <= 122 {
			r[i] = r[i] + shift
			if r[i] > 122 {
				r[i] = r[i] - 26
			}
		}
		if r[i] >= 65 && r[i] <= 90 {
			r[i] = r[i] + shift
			if r[i] > 90 {
				r[i] = r[i] - 26
			}
		}
	}
	s := string(r)
	return s
}