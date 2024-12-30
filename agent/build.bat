@echo off
setlocal enabledelayedexpansion

:: Set Version
set BUILD_VERSION=1.9.1
:: Set environment variables
set AGENT_PACKAGE=github.com/bytedance/Elkeid/agent/agent
set WORKDIR=%cd%

:: Clean up previous builds
rmdir /s /q output
rmdir /s /q build
mkdir output
mkdir build

:: Loop through the architectures (amd64, arm64)
for %%A in (amd64) do (
    set GOARCH=%%A

    :: Build elkeidctl
    cd %WORKDIR%\deploy\control
    go build -o ../../build\elkeidctl.exe
    cd %WORKDIR%

    :: Build elkeid-agent with version
    go build -tags product -ldflags "-X %AGENT_PACKAGE%.Version=%BUILD_VERSION%" -o build\elkeid-agent.exe
    echo binary build done.

    :: Copy the binaries to output
    cd %WORKDIR%
    copy build\*.exe output\
    
    :: Clean up build directory
    rmdir /s /q build
    mkdir build
)

endlocal