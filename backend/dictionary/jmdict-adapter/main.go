package main

import(
	"fmt"
	"os"
	"encoding/json"
	"log"
	"strings"
	"github.com/programmer8760/japanese-parser/backend/parser"
	"github.com/programmer8760/japanese-parser/backend/types"
)

func toString(v any) string {
    if s, ok := v.(string); ok {
        return s
    }
    return ""
}

func toInt(v any) int {
    switch t := v.(type) {
    case float64:
        return int(t)
    case int:
        return t
    default:
        return 0
    }
}
func toStringSlice(v any) []string {
    arr, ok := v.([]any)
    if !ok {
        return nil
    }
    res := make([]string, len(arr))
    for i, el := range arr {
        res[i] = toString(el)
    }
    return res
}

func main() {
	pathRu := "backend/dictionary/jmdict-adapter/" + "jmdict-russian/"

	filesRu, err := os.ReadDir(pathRu)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var dictRu []types.DictionaryEntry

	parser, _ := parser.NewParser()

	for _, file := range filesRu {
		if strings.HasPrefix(file.Name(), "term_bank") {
			fileBytes, err := os.ReadFile(pathRu + file.Name())
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
			}

			var raw [][]any

			json.Unmarshal(fileBytes, &raw)

			for _, row := range raw {
				split, err := parser.Tokenize(toString(row[0]))
				if len(split) == 1 && err == nil && toStringSlice(row[5])[0] != ""  {
					dictRu = append(dictRu, types.DictionaryEntry{
						Kanji: toString(row[0]),
						Reading: toString(row[1]),
						Translations: toStringSlice(row[5]),
						WordID: toInt(row[6]),
					})
				}
			}
		}
	}

	dictRuMarshal, err := json.MarshalIndent(dictRu, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}

//	err = os.WriteFile(pathRu + "merged.json", dictRuMarshal, 0755)
//	if err != nil {
//		fmt.Println("Error writing file:", err)
//	}

	err = os.WriteFile("backend/dictionary/" + "dictionary.json", dictRuMarshal, 0755)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
