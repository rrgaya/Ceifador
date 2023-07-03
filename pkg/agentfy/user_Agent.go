package agentfy

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

	appleModels := []string{
		"iPhone 11",
		"iPhone XR",
		"iPhone X",
		"iPhone 8",
		"iPhone 7",
		"iPhone SE",
	}

	xiaomiModels := []string{
		"Mi 11",
		"Mi 10",
		"Redmi Note 9",
		"Redmi Note 8",
		"Redmi 9",
		"Redmi 8",
	}

	lgModels := []string{
		"LG G8",
		"LG V40",
		"LG G7",
		"LG Q6",
		"LG K40",
		"LG Stylo 5",
	}

	samsungModels := []string{
		"Galaxy S21",
		"Galaxy S20",
		"Galaxy Note 20",
		"Galaxy A71",
		"Galaxy A51",
		"Galaxy J7",
	}

	brandIndex := rand.Intn(4)
	androidVersion := androidVersions[rand.Intn(len(androidVersions))]

	var userAgent string

	switch brandIndex {
	case 0: // Apple
		appleModel := appleModels[rand.Intn(len(appleModels))]
		userAgent = fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/%s Mobile/15E148 Safari/604.1", androidVersion, appleModel)
	case 1: // Xiaomi
		xiaomiModel := xiaomiModels[rand.Intn(len(xiaomiModels))]
		userAgent = fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Mobile Safari/537.36", androidVersion, xiaomiModel, chromeVersions[rand.Intn(len(chromeVersions))])
	case 2: // LG
		lgModel := lgModels[rand.Intn(len(lgModels))]
		userAgent = fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Mobile Safari/537.36", androidVersion, lgModel, chromeVersions[rand.Intn(len(chromeVersions))])
	case 3: // Samsung
		samsungModel := samsungModels[rand.Intn(len(samsungModels))]
		userAgent = fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Mobile Safari/537.36", androidVersion, samsungModel, chromeVersions[rand.Intn(len(chromeVersions))])
	default: // Caso inesperado
		userAgent = ""
	}

	return userAgent
}

// func main() {
// 	userAgent := GenerateRandomUserAgentAndroid()
// 	fmt.Println(userAgent)
// }
