package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "project", "name")
	flag.Parse()
	// 定义要创建的目录结构
	directories := []string{
		fmt.Sprintf("%s/cmd", name),
		fmt.Sprintf("%s/cmd/main", name),
		fmt.Sprintf("%s/internal", name),
		fmt.Sprintf("%s/pkg", name),
		fmt.Sprintf("%s/config", name),
	}

	// 创建每个目录
	for _, dir := range directories {
		err := createDir(dir)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
		} else {
			fmt.Printf("Successfully created directory: %s\n", dir)
		}
	}
	os.WriteFile(fmt.Sprintf("%s/Dockerfile", name), []byte(dockerfile), 0644)
	os.WriteFile(fmt.Sprintf("%s/config/config.yaml", name), []byte(""), 0644)
	os.WriteFile(fmt.Sprintf("%s/start.sh", name), []byte(startSH), 0644)
	os.WriteFile(fmt.Sprintf("%s/cmd/main/main.go", name), []byte(mainFile), 0644)

	if err := runGoModInit(name); err != nil {
		fmt.Printf("Error initializing Go module: %v\n", err)
	}
	if err := runGoModTidy(name); err != nil {
		fmt.Printf("Error initializing Go module: %v\n", err)
	}
}

// createDir 创建目录，兼容不同操作系统
func createDir(dir string) error {
	// 使用 filepath.Join 在不同操作系统中生成正确的路径
	// 这里 dir 已经是使用相对路径，直接使用即可

	// 检查目标目录是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 尝试创建目录
		aserr := os.MkdirAll(dir, os.ModePerm)
		return aserr
	}

	return nil
}
func runGoModInit(name string) error {
	// 设置命令，并通过 exec.Command 运行
	cmd := exec.Command("go", "mod", "init", name)

	// 设置当前工作目录为项目根目录
	cmd.Dir = name

	// 获取命令的输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run command: %w, output: %s", err, string(output))
	}

	fmt.Printf("Successfully initialized Go module: %s\n", string(output))
	return nil
}
func runGoModTidy(name string) error {
	// 设置命令，并通过 exec.Command 运行
	cmd := exec.Command("go", "mod", "tidy", name)

	// 设置当前工作目录为项目根目录
	cmd.Dir = name

	// 获取命令的输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run command: %w, output: %s", err, string(output))
	}

	fmt.Printf("Successfully initialized Go module: %s\n", string(output))
	return nil
}
