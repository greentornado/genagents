package genagents

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var rvs = []string{"11.0", "40.0", "42.0", "43.0", "47.0", "50.0", "52.0", "53.0", "54.0",
	"61.0", "66.0", "67.0"}

var linuxArch = []string{"i686", "x86_64"}

var linuxDistro = []string{"", "; Ubuntu/14.10", "; Ubuntu/16.10", "; Ubuntu/19.10", "; Ubuntu"}

var windowsNT = []string{"5.1", "5.2", "6.0", "6.1", "6.2", "6.3", "6.4", "9.0", "10.0"}

var windowsArch = []string{"; WOW64", "; Win64; x64", "; ARM", ""}

var windowsTrident = []string{"", "; Trident/5.0", "; Trident/6.0", "; Trident/7.0"}

var macV = []string{"Intel Mac OS X 10_12", "Intel Mac OS X 10_11", "Intel Mac OS X 10_10_1",
	"Intel Mac OS X 10_9_3"}

var safariV = []string{"536.25", "536.26", "536.26.17", "536.28.10", "536.29.13", "536.30.1",
	"537.32", "537.36", "537.43.58", "537.73.11", "537.85.17", "537.71",
	"537.73.11", "537.75.14", "537.76.4", "537.77.4", "537.78.2",
	"537.85.17", "538.35.8", "600.6.3", "600.7.12", "601.1.56", "601.2.7",
	"601.3.9", "601.4.4", "601.5.17", "601.6.17", "601.7.1", "601.7.8",
	"602.1.50", "602.2.14", "602.3.12", "602.4.8", "603.1.30", "603.2.4",
	"603.3.8", "604.1.28"}

var edge = []string{"", "", " Edge/16.16299"}

var firefoxV = []string{"38.0", "40.1", "43.0", "50.0", "52.0", "53.0", "60.0", "60.0.2", "60.0.1",
	"61.0", "61.0.1", "66.0", "67.0"}

var prestoP = []string{"2.12.388", "2.12.407", "22.9.168"}

var prestoV = []string{"12.00", "12.14", "12.16"}

var oprV = []string{"43.0.2442.991", "36.0.2130.32", "56.0.3051.52", "47.0.2631.39",
	"42.0.2393.94", "49.0.2711.0", "34.0.2036.25", "52.0.2871.99",
	"33.0.1990.115", "53.0.2907.99"}

var opera = []string{"Opera/9.80", "Opera/12.0"}

func randomSlice(slc []string) string {
	return slc[rand.Intn(len(slc))]
}

func rv() string {
	rand.Seed(time.Now().Unix())
	if rand.Intn(5) < 3 {
		return ""
	}

	return "; rv:" + randomSlice(rvs)
}

func linux() string {
	return "X11; Linux " + randomSlice(linuxArch) + randomSlice(linuxDistro)
}

func windows() string {
	return "Windows " + randomSlice(windowsNT) + randomSlice(windowsArch) + randomSlice(windowsTrident)
}

func mac() string {
	return "Macintosh; " + randomSlice(macV)
}

func appleWebkit() string {
	return "AppleWebKit/" + randomSlice(safariV) + " (KHTML, like Gecko) "
}

func safari() string {
	return " Safari/" + randomSlice(safariV) + randomSlice(edge)
}

func chrome() string {
	return "Chrome/" + randomSlice(chromeV)
}

func firefox() string {
	return "Gecko/20100101 Firefox/" + randomSlice(firefoxV)
}

func presto() string {
	return fmt.Sprintf("Presto/%s Version/%s", randomSlice(prestoP), randomSlice(prestoV))
}

func opr() string {
	return " OPR/" + randomSlice(oprV)
}

func product() string {
	rand.Seed(time.Now().Unix())
	if rand.Intn(10) < 8 {
		return "Mozilla/5.0"
	}

	return randomSlice(opera)
}

func os() string {
	return fmt.Sprintf("(%s%s)", randomSlice([]string{linux(), windows(), mac()}), rv())
}

func browser(os string, prod string) string {
	if strings.Contains(prod, "Opera") {
		if rand.Intn(1) == 0 {
			return presto()
		}
		return appleWebkit() + chrome() + safari() + opr()

	}

	if strings.Contains(os, "X11") && !strings.Contains(os, "rv") {
		return appleWebkit() + chrome() + safari()
	}

	rand.Seed(time.Now().Unix())
	var ri = rand.Intn(100)
	if ri < 5 {
		return "like Gecko"
	} else if ri < 50 {
		return appleWebkit() + chrome() + safari()
	} else {
		return firefox()
	}
}

// GenAgent func
func GenAgent() string {
	prod := product()
	os := os()
	return fmt.Sprintf("%s %s %s", prod, os, browser(os, prod))
}
