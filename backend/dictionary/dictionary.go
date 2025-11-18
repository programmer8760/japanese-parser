package dictionary

import(
	"encoding/json"
	_ "embed"

	"github.com/programmer8760/japanese-parser/backend/types"
)

//go:embed dictionary.json
var dictionaryJSON []byte

type Dictionary struct {
	db map[string][]types.DictionaryEntry
}

func NewDictionary() (*Dictionary, error) {
	var jsonEntries []types.DictionaryEntry
	if err := json.Unmarshal(dictionaryJSON, &jsonEntries); err != nil {
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
