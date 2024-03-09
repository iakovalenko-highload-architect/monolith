package converter

import (
	"github.com/mitchellh/mapstructure"
)

func ConvertMapToStruct(m map[string]interface{}, s interface{}) error {
	return mapstructure.Decode(m, s)
}
