package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	NginxConfTemplate   = "./template/nginx.conf.template"
	RunScriptTemplate   = "./template/run.sh.template"
	StartScriptTemplate = "./template/start.sh.template"
	StopScriptTemplate  = "./template/stop.sh.template"
	MainLuaTemplate     = "./template/main.lua.template"
)

func main() {
	argCount := len(os.Args)
	if argCount < 4 {
		fmt.Println("Argument format: project_name service_name service_port(example: openresty-chat account-service 80)")
		os.Exit(-1)
	}
	projectName := os.Args[1]
	serviceName := os.Args[2]
	servicePort, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	serviceDirectoryPath, confDirectoryPath, srcDirectoryPath := createMustDirectory(projectName, serviceName)

	createRunScript(serviceDirectoryPath, projectName, serviceName, servicePort)
	createStartScript(serviceDirectoryPath, projectName, serviceName, servicePort)
	createStopScript(serviceDirectoryPath, projectName, serviceName, servicePort)

	createNginxConfFile(confDirectoryPath, projectName, serviceName, servicePort)

	createMainLua(srcDirectoryPath, projectName, serviceName, servicePort)

	fmt.Println("finished!")
}

func createRunScript(serviceDirectoryPath string, projectName string, serviceName string, servicePort int) {
	readContent := readFileContent(RunScriptTemplate)
	runScriptContent := replaceTemplateContent(readContent, projectName, serviceName, servicePort)
	runFileName := fmt.Sprintf("%s/run.sh", serviceDirectoryPath)
	writeFileContent(runFileName, runScriptContent)
	fmt.Println(runScriptContent)
	fmt.Println()
}

func createStartScript(serviceDirectoryPath string, projectName string, serviceName string, servicePort int) {
	readContent := readFileContent(StartScriptTemplate)
	startScriptContent := replaceTemplateContent(readContent, projectName, serviceName, servicePort)
	startFileName := fmt.Sprintf("%s/start.sh", serviceDirectoryPath)
	writeFileContent(startFileName, startScriptContent)
	fmt.Println(startScriptContent)
	fmt.Println()
}

func createStopScript(serviceDirectoryPath string, projectName string, serviceName string, servicePort int) {
	readContent := readFileContent(StopScriptTemplate)
	stopScriptContent := replaceTemplateContent(readContent, projectName, serviceName, servicePort)
	stopFileName := fmt.Sprintf("%s/stop.sh", serviceDirectoryPath)
	writeFileContent(stopFileName, stopScriptContent)
	fmt.Println(stopScriptContent)
	fmt.Println()
}

func createNginxConfFile(confDirectoryPath string, projectName string, serviceName string, servicePort int) {
	readContent := readFileContent(NginxConfTemplate)
	nginxConfContent := replaceTemplateContent(readContent, projectName, serviceName, servicePort)
	nginxFileName := fmt.Sprintf("%s/%s-nginx.conf", confDirectoryPath, serviceName)
	writeFileContent(nginxFileName, nginxConfContent)
	fmt.Println(nginxConfContent)
	fmt.Println()
}

func createMainLua(srcDirectoryPath string, projectName string, serviceName string, servicePort int) {
	readContent := readFileContent(MainLuaTemplate)
	mainLuaContent := replaceTemplateContent(readContent, projectName, serviceName, servicePort)
	mainLuaName := fmt.Sprintf("%s/main.lua", srcDirectoryPath)
	writeFileContent(mainLuaName, mainLuaContent)
	fmt.Println(mainLuaContent)
	fmt.Println()
}

func createMustDirectory(projectName string, serviceName string) (serviceDirectoryPath string, confDirectoryPath string, srcDirectoryPath string) {
	serviceDirectoryPath = fmt.Sprintf("./%s/%s", projectName, serviceName)
	confDirectoryPath = fmt.Sprintf("%s/conf", serviceDirectoryPath)
	createDirectory(confDirectoryPath)
	srcDirectoryPath = fmt.Sprintf("%s/src", serviceDirectoryPath)
	createDirectory(srcDirectoryPath)
	fmt.Println()
	return
}

func createDirectory(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("createDirectory:", path)
}

func readFileContent(filename string) string {
	reader, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("err: ", err)
		return ""
	}
	result := string(reader)
	return result
}

func writeFileContent(filename string, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("createFile:", filename)
}

func replaceTemplateContent(templateContent string, projectName string, serviceName string, servicePort int) string {
	result := templateContent
	result = strings.Replace(result, "${service_name}", serviceName, -1)
	result = strings.Replace(result, "${project_name}", projectName, -1)
	result = strings.Replace(result, "${service_port}", strconv.Itoa(servicePort), -1)
	result = strings.Replace(result, "\r\n", "\n", -1)
	return result
}
