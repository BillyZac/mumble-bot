package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var baseConfig = Config{
	TwitterConsumerKey:       "consumerkey",
	TwitterConsumerSecret:    "consumersecret",
	TwitterAccessToken:       "accesstoken",
	TwitterAccessTokenSecret: "accesstokensecret",
}

var configVariations = []struct {
	field string
	value reflect.Value
	valid bool
}{
	{"TwitterConsumerKey", reflect.ValueOf("consumerkey"), true},
	{"TwitterConsumerKey", reflect.ValueOf(""), false},
	{"TwitterConsumerSecret", reflect.ValueOf("consumersecret"), true},
	{"TwitterConsumerSecret", reflect.ValueOf(""), false},
	{"TwitterAccessToken", reflect.ValueOf("accesstoken"), true},
	{"TwitterAccessToken", reflect.ValueOf(""), false},
	{"TwitterAccessTokenSecret", reflect.ValueOf("accesstokensecret"), true},
	{"TwitterAccessTokenSecret", reflect.ValueOf(""), false},
}

func TestConfigFieldValidations(t *testing.T) {
	for _, test := range configVariations {
		config := baseConfig
		reflect.ValueOf(&config).Elem().FieldByName(test.field).Set(test.value)
		err := config.validate()

		if test.valid {
			msg := fmt.Sprintf("config should validate with valid data, (field: \"%s\", value:\"%+v\")", test.field, test.value)
			assert.NoError(t, err, msg)
		} else {
			msg := fmt.Sprintf("config shouldn't validate with invalid data, (field: \"%s\", value:\"%+v\")", test.field, test.value)
			assert.Error(t, err, msg)
		}
	}
}
