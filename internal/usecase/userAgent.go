package usecase

import (
	"fmt"
	"math/rand"
)

func GenerateRandomUserAgentAndroid() string {
	androidVersions := []string{
		"4.0.3",
		"4.1.1",
		"4.2.2",
		"4.3",
		"4.4",
		"5.0",
		"5.1",
		"6.0",
		"7.0",
		"7.1",
		"8.0",
		"8.1",
		"9.0",
		"10.0",
		"11.0",
		"12.0",
	}

	chromeVersions := []string{
		"86.0.4240.198",
		"87.0.4280.141",
		"88.0.4324.150",
		"89.0.4389.82",
		"90.0.4430.72",
		"91.0.4472.124",
		"92.0.4515.159",
	}

	firefoxVersions := []string{
		"78.0",
		"79.0",
		"80.0",
		"81.0",
		"82.0",
		"83.0",
		"84.0",
		"85.0",
		"86.0",
		"87.0",
		"88.0",
		"89.0",
		"90.0",
		"91.0",
		"92.0",
	}

	androidVersion := androidVersions[rand.Intn(len(androidVersions))]

	chromeUserAgent := fmt.Sprintf("Mozilla/5.0 (Linux; Android %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Mobile Safari/537.36", androidVersion, chromeVersions[rand.Intn(len(chromeVersions))])
	firefoxUserAgent := fmt.Sprintf("Mozilla/5.0 (Android %s; Mobile; rv:%s) Gecko/%s Firefox/%s", androidVersion, androidVersion, firefoxVersions[rand.Intn(len(firefoxVersions))], firefoxVersions[rand.Intn(len(firefoxVersions))])

	if rand.Intn(2) == 0 {
		return chromeUserAgent
	}
	return firefoxUserAgent
}
