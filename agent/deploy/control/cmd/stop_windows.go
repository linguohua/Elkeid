/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/nightlyone/lockfile"
)

func sysvinitStop() error {
	os.RemoveAll(crontabFile)
	file, err := lockfile.New(agentPidFile)
	if err != nil {
		return err
	}
	p, err := file.GetOwner()
	if err == nil {
		var pids []int
		pids, err := GetProcs(p.Pid)
		if err != nil {
			return err
		}
		for _, pid := range pids {
			killProcess(pid)
		}
		ticker := time.NewTicker(time.Millisecond * time.Duration(100))
		defer ticker.Stop()
		timeout := time.NewTimer(time.Second * time.Duration(30))
		i := 0
		defer timeout.Stop()
		for {
			select {
			case <-ticker.C:
				pids = CheckPids(pids)
				if len(pids) == 0 {
					return nil
				}
				if i%50 == 0 {
					fmt.Printf("wait %v subprocess to exit...\n", len(pids))
				}
				i++
			case <-timeout.C:
				fmt.Fprintln(os.Stderr, "stop timeout, will kill all subprocess...")
				for _, pid := range pids {
					killProcess(pid)
				}
				return nil
			}
		}
	}
	return nil
}

func killProcess(pid int) {
	processHandle, err := syscall.OpenProcess(syscall.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		fmt.Println("Error opening process:", err)
		return
	}
	defer syscall.CloseHandle(processHandle)

	// Terminate the process
	err = syscall.TerminateProcess(processHandle, 0)
	if err != nil {
		fmt.Println("Error terminating process:", err)
		return
	}
}
