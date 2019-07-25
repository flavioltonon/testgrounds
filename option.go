package testgrounds

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

const OPTION = "option"

type Option interface {
	Type() string
}

var options = make(OptionsMap, 0)

func RegisterOption(option Option) {
	if _, exists := options[option.Type()]; exists {
		panic(fmt.Sprintf("option %s has already been registered", option.Type()))
	}

	options[option.Type()] = option
}

func RegisteredOptions() OptionsMap {
	return options
}

func NewOption(t string) (Option, error) {
	if _, exists := options[t]; !exists {
		return nil, fmt.Errorf("'%s' is not a registered option type", t)
	}
	return options[t], nil
}

type OptionsMap map[string]Option

func (opts Options) ByType(t string) Options {
	var optionsByType = make(Options, 0)

	for _, opt := range opts {
		if opt.Type() == t {
			optionsByType = append(optionsByType, opt)
		}
	}

	return optionsByType
}

type Options []Option

func (opts Options) Merge() (Option, error) {
	var mainType = opts[0].Type()

	var o = make(map[string]interface{}, 0)

	for _, opt := range opts {
		if opt.Type() != mainType {
			return nil, errors.New("merging options is only available for arrays of options of the same type")
		}

		value := reflect.ValueOf(opt)
		if value.Kind() != reflect.Struct {
			return nil, fmt.Errorf("unexpected object type: %v", value.Kind())
		}

		for i := 0; i < value.Type().NumField(); i++ {
			if value.Field(i).Interface() != nil {
				o[value.Type().Field(i).Name] = value.Field(i).String()
			}
		}
	}

	bytes, err := json.Marshal(o)
	if err != nil {
		return nil, errors.New("failed to serialize bucket option")
	}

	option, err := NewOption(mainType)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &option)
	if err != nil {
		return nil, errors.New("failed to decode bucket option")
	}

	return option, nil
}
