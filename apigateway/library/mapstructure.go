package library

import (
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/s3f4/mu/log"
)

// toTimeHook to time
func toTimeHook() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}

// DecodeMap mapstructure
func DecodeMap(source, result interface{}) error {
	cfg := &mapstructure.DecoderConfig{
		Metadata:   nil,
		Result:     result,
		TagName:    "json",
		DecodeHook: mapstructure.ComposeDecodeHookFunc(toTimeHook()),
	}

	decoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		log.Errorf("mapstructrure.decode.NewDecoder", err)
		return err
	}

	if err := decoder.Decode(source); err != nil {
		log.Errorf("mapstructrure.decode.Decode", err)
		return err
	}

	return nil
}
