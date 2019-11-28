package logic

var pastFourScores = [...]float64{5.0, 4.0, 2.0, 1.0}

func GetSize() float64 {
	oldScores := pastFourScores[0] + pastFourScores[1]
	newScores := pastFourScores[2] + pastFourScores[3]
	diff := newScores - oldScores
	if diff > 0 {
		size := 600 + diff * 60
		if size < 2000 {
			return size
		}
		return 2000
	}
	if diff > -5 && diff <= 0 {
		return 100 + 18 * diff
	}
	return 10
}

func SetScore(val float64) bool {
	for i := range pastFourScores {
		if i < len(pastFourScores) - 1 {
			pastFourScores[i] = pastFourScores[i + 1]
		}
	}
	pastFourScores[3] = val
	return true
}