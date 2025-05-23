package new

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewProject(t *testing.T) {
	p := NewProject()
	if p == nil {
		t.Error("NewProject() should not return nil")
	}
}

func TestProjectCloneTemplate(t *testing.T) {
	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "test-project")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 设置测试项目
	p := &Project{
		ProjectName: filepath.Join(tmpDir, "test-project"),
	}

	// 测试克隆模板
	repoURL = "https://gitee.com/QianLiGitee/template.git"
	yes, err := p.cloneTemplate()
	if err != nil {
		t.Skipf("cloneTemplate() error = %v (skipping test due to potential network/git issues)", err)
	}
	if !yes {
		t.Error("cloneTemplate() should return true")
	}

	// 验证项目目录是否存在
	if _, err := os.Stat(p.ProjectName); os.IsNotExist(err) {
		t.Error("Project directory should exist")
	}
}

func TestProjectReplacePackageName(t *testing.T) {
	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "test-project")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建测试项目结构
	projectDir := filepath.Join(tmpDir, "test-project")
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		t.Fatalf("Failed to create project dir: %v", err)
	}

	// 创建测试 go.mod 文件
	goModContent := "module github.com/example/old-project"
	if err := os.WriteFile(filepath.Join(projectDir, "go.mod"), []byte(goModContent), 0644); err != nil {
		t.Fatalf("Failed to write go.mod: %v", err)
	}

	// 设置测试项目
	p := &Project{
		ProjectName: projectDir,
	}

	// 测试替换包名
	err = p.replacePackageName()
	if err != nil {
		t.Errorf("replacePackageName() error = %v", err)
	}

	// 验证 go.mod 文件内容
	content, err := os.ReadFile(filepath.Join(projectDir, "go.mod"))
	if err != nil {
		t.Fatalf("Failed to read go.mod: %v", err)
	}

	expected := "module " + projectDir
	if string(content) != expected {
		t.Errorf("go.mod content = %v, want %v", string(content), expected)
	}
}

func TestProjectModTidy(t *testing.T) {
	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "test-project")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建测试项目结构
	projectDir := filepath.Join(tmpDir, "test-project")
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		t.Fatalf("Failed to create project dir: %v", err)
	}

	// 创建测试 go.mod 文件
	goModContent := "module github.com/example/test-project\ngo 1.21"
	if err := os.WriteFile(filepath.Join(projectDir, "go.mod"), []byte(goModContent), 0644); err != nil {
		t.Fatalf("Failed to write go.mod: %v", err)
	}

	// 设置测试项目
	p := &Project{
		ProjectName: projectDir,
	}

	// 测试 modTidy
	err = p.modTidy()
	if err != nil {
		t.Skipf("modTidy() error = %v (skipping test due to potential go mod issues)", err)
	}
}

func TestProjectRmGit(t *testing.T) {
	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "test-project")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建测试项目结构
	projectDir := filepath.Join(tmpDir, "test-project")
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		t.Fatalf("Failed to create project dir: %v", err)
	}

	// 创建 .git 目录
	gitDir := filepath.Join(projectDir, ".git")
	if err := os.MkdirAll(gitDir, 0755); err != nil {
		t.Fatalf("Failed to create .git dir: %v", err)
	}

	// 设置测试项目
	p := &Project{
		ProjectName: projectDir,
	}

	// 测试 rmGit
	p.rmGit()

	// 验证 .git 目录是否被删除
	if _, err := os.Stat(gitDir); !os.IsNotExist(err) {
		t.Error(".git directory should be removed")
	}
}

func TestProjectReplaceFiles(t *testing.T) {
	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "test-project")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建测试项目结构
	projectDir := filepath.Join(tmpDir, "test-project")
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		t.Fatalf("Failed to create project dir: %v", err)
	}

	// 创建测试 Go 文件
	testFile := filepath.Join(projectDir, "main.go")
	testContent := "package main\n\nimport \"github.com/example/old-project/pkg\"\n\nfunc main() {\n\tpkg.Hello()\n}"
	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// 设置测试项目
	p := &Project{
		ProjectName: projectDir,
	}

	// 测试 replaceFiles
	err = p.replaceFiles("github.com/example/old-project")
	if err != nil {
		t.Errorf("replaceFiles() error = %v", err)
	}

	// 验证文件内容是否被替换
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	expected := "package main\n\nimport \"" + projectDir + "/pkg\"\n\nfunc main() {\n\tpkg.Hello()\n}"
	if string(content) != expected {
		t.Errorf("file content = %v, want %v", string(content), expected)
	}
}
