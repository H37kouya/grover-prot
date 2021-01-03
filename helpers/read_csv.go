package helpers

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type xyAxis struct {
	X float64
	Y float64
}

func ReadCsv(path string, count int) ([]xyAxis, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	ret := make([]xyAxis, 0, count)

	for i := 0; i < count; i++ {
		record, err := r.Read()
		if i == 0 {
			i++
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		floatVal, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return nil, err
		}
		iFloatVal := float64(i)

		ret = append(ret, xyAxis{
			X: iFloatVal,
			Y: floatVal,
		})
	}

	return ret, nil
}
