package semver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	semverMajor = "major"
	semverMinor = "minor"
	semverPatch = "patch"
)

func Upversion(step, current string) (string, error) {
	if err := validateStep(step); err != nil {
		return "", err
	}
	prefix := ""
	if len(current) > 0 && current[0] == 'v' {
		prefix = "v"
		current = current[1:]
	}
	major, minor, patch, err := validateAndParseSemver(current)
	if err != nil {
		return "", err
	}

	if step == semverMajor {
		major++
		minor = 0
		patch = 0
	}
	if step == semverMinor {
		minor++
		patch = 0
	}
	if step == semverPatch {
		patch++
	}
	return fmt.Sprintf("%s%d.%d.%d", prefix, major, minor, patch), nil
}

func validateStep(step string) error {
	if step == "" {
		return errors.New("no step value given")
	}
	if step != semverMajor && step != semverMinor && step != semverPatch {
		return fmt.Errorf("invalid step value '%s', valid values are ['major', 'minor', 'patch']", step)
	}
	return nil
}

func validateAndParseSemver(current string) (int, int, int, error) {
	split := strings.Split(current, ".")
	humanErr := fmt.Errorf("current version (%s) is not semver", current)
	if len(split) != 3 {
		return 0, 0, 0, humanErr
	}

	majorStr, minorStr, patchStr := split[0], split[1], split[2]

	major, err := strconv.Atoi(majorStr)
	if err != nil {
		return 0, 0, 0, humanErr
	}
	minor, err := strconv.Atoi(minorStr)
	if err != nil {
		return 0, 0, 0, humanErr
	}
	patch, err := strconv.Atoi(patchStr)
	if err != nil {
		return 0, 0, 0, humanErr
	}
	return major, minor, patch, nil

}
