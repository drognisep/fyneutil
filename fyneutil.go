package util

func maxF32(base, other float32) float32 {
	if other > base {
		return other
	}
	return base
}

func minF32(base, other float32) float32 {
	if other < base {
		return other
	}
	return base
}
