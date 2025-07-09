package winreg

import (
	"fmt"
	"strings"
)

func ParseRegistryHiveName(hive string) (string, error) {

	switch hive {
	case KeyLocalMachine, KeyCurrentUser, KeyClassesRoot, KeyUsers, KeyCurrentConfig:
		return hive, nil
	}

	switch strings.ToLower(hive) {
	case "hklm":
		return KeyLocalMachine, nil
	case "hkcu":
		return KeyCurrentUser, nil
	case "hkcr":
		return KeyClassesRoot, nil
	case "hku":
		return KeyUsers, nil
	case "hkcc":
		return KeyCurrentConfig, nil
	}

	return "", fmt.Errorf("unknown hive: %s", hive)
}
