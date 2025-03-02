package parsinglogfiles

import "regexp"

// import "fmt"

// Search log with starting strings like [TRC] or [DBG] or [INF] or [WRN] or [ERR] or [FTL]
func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	if err != nil {
		panic(err.Error())
	}
	matchStr := re.Find([]byte(text))
	// fmt.Printf("text:%s\n", text)
	// fmt.Printf("matchStr:%s\n", string(matchStr))
	if len(matchStr) != 0 {
		return true
	}
	return false
}

// Log lines have less than (<) and more than (>).
// Between two letters, ~ or * or = or - or "" exist.
// exp) <-->, <==>, <>, <--=*=-->
func SplitLogLine(text string) []string {
	// fmt.Println("text:" + text)
	re, err := regexp.Compile(`<(~|\*|=|\-)*>`)
	if err != nil {
		panic(err.Error())
	}
	retArr := re.Split(text, -1)
	// fmt.Printf("retArr: %v\n", retArr)
	return retArr
}

// password includes upper letters and lower letters
// There is quotation mark (") before password.
// There is quotation mark (") after password.
func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`\".*(?i)password(?-i).*\"`)
	if err != nil {
		panic(err.Error())
	}
	count := 0
	for _, v := range lines {
		matchStr := re.FindString(v)
		if matchStr != "" {
			count++
		}
	}
	return count
}

// text includes missing string like end-of-line12345
// The numbers are something, not fixed.
func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end\-of\-line\d+`)
	if err != nil {
		panic(err.Error())
	}
	return re.ReplaceAllString(text, "")
}

// If log has user name, "User " and user name.
// User name is alphabets and numbers.
// The space count is more than equal one after user name.
func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`User +([a-zA-Z0-9]+)`)
	if err != nil {
		panic(err.Error())
	}

	returnStrArray := []string{}
	replaceStr := ""
	for _, v := range lines {
		// catch user name
		matchStrArr := re.FindStringSubmatch(v)
		// fmt.Printf("matchStrArr: %v\n", matchStrArr)
		matchStr := ""
		if len(matchStrArr) > 1 {
			// change line which user name exists
			matchStr = matchStrArr[1]
			// fmt.Printf("matchStr:%s\n", matchStr)
			// add [USR] and user name and space at the first of line
			replaceStr = "[USR] " + matchStr + " " + v
		} else {
			// not change
			replaceStr = v
		}
		// fmt.Printf("replaceStr:%s\n", replaceStr)
		returnStrArray = append(returnStrArray, replaceStr)
	}
	return returnStrArray
}
