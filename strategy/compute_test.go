package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriceCompute_Do(t *testing.T) {
	computeValues := []struct {
		price         float64
		promotionType string
		result        float64
	}{
		{200, "满200减100", 100},
		{150, "满200减100", 150},
		{200.3583654875, "满200减 100.3235124545", 100.03},
		{200.3593654875, "满200减 100.3235124545", 100.04},
		{1, "满1减0.8", 0.20},
		{200, "8折", 160},
		{200, "8.8折", 176},
	}

	for _, v := range computeValues {
		com := &PriceCompute{v.price, v.promotionType, nil}
		res := com.Do()
		assert.Exactly(t, v.result, res, "计算不正确")
	}
}
