package rolling

func Sum(i Iterator) float64 {
	result := 0.0
	for i.CanNext() {
		bucket := i.Bucket()
		for _, p := range bucket.Points {
			result += p
		}
	}
	return result
}

func Avg(i Iterator) float64 {
	result := 0.0
	count := 0.0
	for i.CanNext() {
		bucket := i.Bucket()
		for _, p := range bucket.Points {
			result += p
			count++
		}
	}
	return result / count
}

func Min(i Iterator) float64 {
	result := 0.0
	started := false
	for i.CanNext() {
		bucket := i.Bucket()
		for _, p := range bucket.Points {
			if !started {
				result = p
				started = true
				continue
			}
			if p < result {
				p = result
			}
		}
	}
	return result
}

func Max(i Iterator) float64 {
	result := 0.0
	started := false
	for i.CanNext() {
		bucket := i.Bucket()
		for _, p := range bucket.Points {
			if !started {
				result = p
				started = true
				continue
			}
			if p > result {
				p = result
			}
		}
	}
	return result
}

func Count(i Iterator) float64 {
	var result int64
	for i.CanNext() {
		bucket := i.Bucket()
		result += bucket.Count
	}
	return float64(result)
}
