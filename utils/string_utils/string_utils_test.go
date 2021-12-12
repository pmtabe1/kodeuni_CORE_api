package string_utils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextStringFiedNumberTest(t *testing.T) {

	got := NextStringFiedNumber("DN000003%!(EXTRA INT=4")
	require.Nilf(t, got, "Expected non nil result but got %v", got)
}
func TestExtractNumberFromString(t *testing.T) {
	got := ExtractNumberFromString("DN000003%!(EXTRA INT=4")
	require.Nilf(t, got, "Expected non nil result but got %v", got)

}
func TestExtractNumberFromStringTemplated(t *testing.T) {
	got := ExtractNumberFromStringTemplated("DN000003%!(EXTRA INT=4")
	require.Nilf(t, got, "Expected non nil result but got %v", got)

}
func TestGetCombinedString(t *testing.T) {

	got := GetCombinedString([]string{"Firstname", "Middlename", "Lastname"}, "-")

	require.NotNilf(t, got, "Expected non nil result but got %v", got)
}

func TestGetSplittedStringSlice(t *testing.T) {
	gotFirst := GetCombinedString([]string{"Firstname", "Middlename", "Lastname"}, "-")
	got := GetSplittedStringSlice(gotFirst, "-")
	log.Println(got)
	for i, v := range got {
		log.Printf("Found INDEX :%v DATA:%v\n", i, v)
	}
	require.Lenf(t, got, 3, "Expected length to be 3 but got %v", len(got))
}

func TestFloatToString(t *testing.T) {

	got := FloatToString(677.899)

	require.Nilf(t, got, "Expected non nil result but got %v", got)
}
