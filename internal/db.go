package internal

import (
	"encoding/json"
	"fmt"
	"strings"

	bolt "go.etcd.io/bbolt"
)

func GetCVE(dbPath, id string) (map[string]interface{}, error) {
	var vuln map[string]interface{}

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("vulnerability"))
		if b == nil {
			return fmt.Errorf("bucket 'vulnerability' introuvable")
		}

		v := b.Get([]byte(id))
		if v == nil {
			return fmt.Errorf("CVE %s non trouvée", id)
		}

		return json.Unmarshal(v, &vuln)
	})

	return vuln, err
}

func SearchByKeyword(dbPath, keyword string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("vulnerability"))
		if b == nil {
			return fmt.Errorf("bucket 'vulnerability' introuvable")
		}

		return b.ForEach(func(k, v []byte) error {
			var vuln map[string]interface{}
			if err := json.Unmarshal(v, &vuln); err != nil {
				return nil
			}

			for _, val := range vuln {
				if str, ok := val.(string); ok && strings.Contains(strings.ToLower(str), strings.ToLower(keyword)) {
					vuln["CVE-ID"] = string(k)
					results = append(results, vuln)
					break
				}
			}

			return nil
		})
	})

	return results, err
}
