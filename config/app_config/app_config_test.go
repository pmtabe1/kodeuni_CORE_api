package app_config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadJSONConfig(t *testing.T) {
	os.Setenv("SAGE_CONFIG", "/etc/integrations/conf/subscription/config.%v.json")
	got := New().LoadJSONConfig("")
	require.NotEmptyf(t, got, "Expected non empty but got %v instead", got)
}

func TestReadConfiguration(t *testing.T) {
	got := New().ReadConfiguration()
	require.NotNilf(t, got, "Expected non nil but got %v instead", got)

}

func TestGetFilepathTemplate(t *testing.T) {

	got := GetFilepathTemplate()
	require.NotEmptyf(t, got, "Expected non empty but got %v instead", got)

}

func TestFromJSON(t *testing.T) {
	jsonString := New().LoadJSONConfig("")
	got := New().FromJSON(jsonString)
	require.NotEmptyf(t, got, "Expected non empty but got %v instead", got)

}
