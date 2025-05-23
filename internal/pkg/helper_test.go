package pkg

import (
	"os"
	"testing"
)

func TestGetProjectName(t *testing.T) {
	// 创建一个临时目录
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 在临时目录中创建一个 go.mod 文件
	goModContent := "module github.com/liqianbro/Ginit"
	if err := os.WriteFile(tmpDir+"/go.mod", []byte(goModContent), 0644); err != nil {
		t.Fatalf("Failed to write go.mod: %v", err)
	}

	// 调用 GetProjectName
	projectName := GetProjectName(tmpDir)
	expected := "github.com/liqianbro/Ginit"
	if projectName != expected {
		t.Errorf("GetProjectName() = %v, want %v", projectName, expected)
	}
}
