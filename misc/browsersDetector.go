/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package misc

import (
	"golang.org/x/sys/windows/registry"
	"runtime"
	"strings"
)

func EdgeDetector() bool {
	if runtime.GOOS != "windows" {
		return true
	} else {
		k, err := registry.OpenKey(registry.CLASSES_ROOT, `Local Settings\Software\Microsoft\Windows\CurrentVersion\AppModel\PackageRepository\Packages`, registry.READ)
		if err != nil {
			return false
		}

		subkeys, _ := k.ReadSubKeyNames(-1)

		for _, subkey := range subkeys {
			if strings.HasPrefix(subkey, "Microsoft.MicrosoftEdge_") {
				return true
			}
		}

		return false
	}
}
