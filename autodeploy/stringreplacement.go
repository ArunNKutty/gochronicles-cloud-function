package autodeploy

import (
	"fmt"
	"net/http"
	"strings"
)

var substrings = []string{"Google","Oracle","Microsoft","Amazon","google","oracle","microsoft","amazon","Adobe"}

//CopyRight adds a copyright symbol if the above strings are encountered
func CopyRightWord(w http.ResponseWriter, r *http.Request) {
	queryParamDisplayHandler(w, r)
}


func checkSubstrings(inputString string) string {
	for _, sub := range substrings {
		if strings.Contains(inputString, sub) {
			inputString = strings.ReplaceAll(inputString,sub,sub+"©")
		}
	}
	return inputString
}

func queryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
	inputSentence := req.FormValue("sentence")
	if inputSentence != "" {
		fmt.Fprint(res, "New Copyrighted Sentence: "+checkSubstrings(inputSentence))
	} else {
		fmt.Fprint(res, "New Copyrighted Sentence: "+"The Input sentence is empty")
	}
}

