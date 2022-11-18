// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/rs/zerolog"

	client "fybrik.io/openmetadata-connector/datacatalog-go-client"
	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
)

const EmptyString = ""

func AppendStrings(a, b string) string {
	if strings.Contains(b, ".") {
		return a + ".\"" + b + "\""
	}
	return a + "." + b
}

// If the tag is in the GenericTags tag category (e.g. "GenericTags.PII"), remove the
// "GenericTags." prefix. Otherwise, do nothing
func StripTag(tag string) string {
	return strings.TrimPrefix(tag, "GenericTags.")
}

func UpdateCustomProperty(customProperties, orig map[string]interface{}, key string, value *string) {
	// update customProperties only if there is a new value
	if value != nil && *value != "" {
		customProperties[key] = *value
		return
	}

	// if there is no new value, revert to original value
	if v, ok := orig[key]; ok && v != "" {
		customProperties[key] = v
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// create random string of letters of length n
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[r.Int64()]
	}
	return string(b)
}

func ExtractColumns(resourceColumns []models.ResourceColumn) []client.Column {
	var ret []client.Column
	for _, rc := range resourceColumns {
		column := *client.NewColumn("STRING", rc.Name)
		column.SetDataLength(0)
		ret = append(ret, column)
	}
	return ret
}

func InterfaceToMap(i interface{}, logger *zerolog.Logger) (map[string]interface{}, bool) {
	m, ok := i.(map[string]interface{})
	if !ok {
		logger.Warn().Msg(fmt.Sprintf("Cannot cast %T to map[string]interface{}", i))
	}
	return m, ok
}

func InterfaceToArray(i interface{}, logger *zerolog.Logger) ([]interface{}, bool) {
	m, ok := i.([]interface{})
	if !ok {
		logger.Warn().Msg(fmt.Sprintf("Cannot cast %T to []interface{}", i))
	}
	return m, ok
}

func GetEnvironmentVariables() (bool, string, string, string) {
	endpoint := os.Getenv("OPENMETADATA_ENDPOINT")
	user := os.Getenv("OPENMETADATA_USER")
	password := os.Getenv("OPENMETADATA_PASSWORD")

	if endpoint == EmptyString || user == EmptyString || password == EmptyString {
		return false, EmptyString, EmptyString, EmptyString
	}
	return true, endpoint, user, password
}
