package utils

import "testing"

func TestGenerateSnowflakeId(t *testing.T) {
	t.Log(GenerateSnowflakeId())
}
