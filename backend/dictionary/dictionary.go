package dictionary

import(
	"os"
	"encoding/json"

	"github.com/programmer8760/japanese-parser/backend/types"
)

type Dictionary struct {
	db map[string][]types.DictionaryEntry
}

func NewDictionary() (*Dictionary, error) {
	fileBytes, err := os.ReadFile("backend/dictionary/dictionary.json")
	if err != nil {
		return nil, err
	}

	var jsonEntries []types.DictionaryEntry
	if err := json.Unmarshal(fileBytes, &jsonEntries); err != nil {
		return nil, err
	}

	db := make(map[string][]types.DictionaryEntry, len(jsonEntries))
	for _, e := range jsonEntries {
		kanji := e.Kanji
		e.Kanji = ""
		db[kanji] = append(db[kanji], e)
	}

	return &Dictionary{db: db}, nil
}

func (d *Dictionary) Lookup(kanji string, reading string) (results []types.DictionaryEntry) {
	if reading != "" {
		for _, e := range d.db[kanji] {
			if e.Reading == reading {
				results = append(results, e)
			}
		}
	} else {
		results = d.db[kanji]
	}

	return
}
