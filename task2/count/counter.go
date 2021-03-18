package count

func Counter (a [][]uint8) (ways uint64) {

	N := len(a)
	M := len(a[0])

	b := make([][]uint64, N)
	for i := range b {
		b[i] = make([]uint64, M)
	}

	b[N-1][0] = 1
	for i:=1; i<M; i++ {
		if a[N-1][i] == 1 {
			b[N-1][i] = b[N-1][i-1]
		} else {
			b[N-1][i] = 0
		}
	}
	for i:=N-2; i>=0; i-- {
		if a[i][0] == 1 {
			b[i][0]=b[i+1][0]
		} else {
			b[i][0] = 0
		}
	}
	for i:=N-2; i>=0; i-- {
		for j:=1; j<M; j++ {
			if a[i][j] == 1 {
				b[i][j] = b[i+1][j]+b[i][j-1]
			} else {
				b[i][j] = 0
			}
		}
	}
	return b[0][M-1]
}
