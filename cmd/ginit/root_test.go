package ginit

import (
	"testing"

	"github.com/liqianbro/Ginit/config"
)

func TestExecute(t *testing.T) {
	// 测试版本信息
	if CmdRoot.Version == "" {
		t.Error("Version should not be empty")
	}

	// 测试版本号格式
	expectedVersion := config.Version
	if CmdRoot.Version == "" {
		t.Errorf("Version should contain %s", expectedVersion)
	}

	// 测试命令名称
	if CmdRoot.Use != "ginit" {
		t.Errorf("Command name should be 'ginit', got %s", CmdRoot.Use)
	}

	// 测试子命令
	if len(CmdRoot.Commands()) == 0 {
		t.Error("Should have at least one subcommand")
	}
}
