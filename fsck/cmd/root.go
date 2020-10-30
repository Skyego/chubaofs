// Copyright 2020 The Chubao Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var c = &cobra.Command{
		Use:   path.Base(os.Args[0]),
		Short: "ChubaoFS fsck tool",
		Args:  cobra.MinimumNArgs(0),
	}

	c.AddCommand(
		newCheckCmd(),
		newCleanCmd(),
	)

	c.PersistentFlags().StringVarP(&MasterAddr, "master", "m", "", "master addresses")
	c.PersistentFlags().StringVarP(&VolName, "vol", "v", "", "volume name")
	c.PersistentFlags().StringVarP(&InodesFile, "inode-list", "i", "", "inode list file")
	c.PersistentFlags().StringVarP(&DensFile, "dentry-list", "d", "", "dentry list file")
	c.PersistentFlags().StringVarP(&MetaPort, "mport", "", "", "prof port of metanode")
	c.PersistentFlags().StringVarP(&VolsFile, "Vols-list", "f", "", "vol list file")
	c.PersistentFlags().IntVarP(&VolMaxBath, "bath", "b", 20, "Vol Max Bath")
	return c
}
