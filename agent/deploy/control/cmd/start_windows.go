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
	"os/exec"
	"syscall"

	"github.com/spf13/viper"
)

func sysvinitStart() error {
	var err error
	// set cgroup
	cgroup, err := NewCGroup(serviceName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create cgroup(even named): %v\n", err.Error())
	}
	cmd := exec.Command(agentFile)
	cmd.Dir = agentWorkDir
	cmd.SysProcAttr = &syscall.SysProcAttr{
		//Setpgid: true, // TODO: windows
	}
	for k, v := range viper.AllSettings() {
		cmd.Env = append(cmd.Env, k+"="+v.(string))
	}
	cmd.Env = append(cmd.Env, "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin")
	err = cmd.Start()
	if err != nil {
		return err
	}
	// maybe no-cgroup
	if cgroup != nil {
		err = cgroup.AddProc(cmd.Process.Pid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to add proc to cgroup: %v\n", err.Error())
		}
	}
	return nil
}
