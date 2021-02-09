package logic

var pastFourScore = []float64{5.0, 4.0, 2.0, 1.0}

// GetSize returns size by considering last 4 scores
func GetSize() float64 {
	oldScore := pastFourScore[0] + pastFourScore[1]
	newScore := pastFourScore[2] + pastFourScore[3]

	diff := newScore - oldScore

	if diff > 0.0 {
		size := 600.0 + diff*60.0
		if size < 2000.0 {
			return size
		}
		return 2000.0
	}

	if diff > -5.0 && diff <= 0.0 {
		return 100.0 + diff*18.0
	}

	return 10.0
}

// SetScore updates past four score
func SetScore(score float64) bool {
	pastFourScore = append(pastFourScore, score)
	pastFourScore = pastFourScore[1:]
	return true
}
