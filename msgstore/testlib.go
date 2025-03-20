package msgstore

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/karelbilek/amqp-test-server/amqp"
	"reflect"
)

// ids to map
func idsToMap() {

}

// Check if the msg store is composed of exactly these keys
func assertKeys(dbName string, keys map[int64]bool) error {
	// Open DB
	db, err := bolt.Open(dbName, 0600, nil)
	defer db.Close()
	if err != nil {
		return err
	}
	err = db.View(func(tx *bolt.Tx) error {
		// Check index
		bucket := tx.Bucket([]byte("message_index"))
		if bucket == nil {
			return nil
		}

		// get from db
		var indexKeys, err1 = keysForBucket(tx, MESSAGE_INDEX_BUCKET)
		var contentKeys, err2 = keysForBucket(tx, MESSAGE_CONTENT_BUCKET)
		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}

		// Check equality
		// TODO: return key diff
		indexNotKeys := subtract(indexKeys, keys)
		keysNotIndex := subtract(keys, indexKeys)
		contentNotKeys := subtract(contentKeys, keys)
		keysNotContent := subtract(keys, contentKeys)

		if !reflect.DeepEqual(keys, indexKeys) {
			return fmt.Errorf("Different values in index!\nindexNotKeys:%q\nkeysNotIndex:%q", indexNotKeys, keysNotIndex)
		}
		if !reflect.DeepEqual(keys, contentKeys) {
			return fmt.Errorf("Different values in content!\ncontentNotKeys:%q\nkeysNotContent:%q", contentNotKeys, keysNotContent)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func subtract(original map[int64]bool, subtractThis map[int64]bool) []int64 {
	ret := make([]int64, 0)
	for id, _ := range original {
		_, found := subtractThis[id]
		if !found {
			ret = append(ret, id)
		}
	}
	return ret
}

func keysForBucket(tx *bolt.Tx, bucketName []byte) (map[int64]bool, error) {
	// Check index
	bucket := tx.Bucket(bucketName)
	if bucket == nil {
		return nil, fmt.Errorf("No bucket!")
	}
	var cursor = bucket.Cursor()
	var keys = make(map[int64]bool)
	for bid, _ := cursor.First(); bid != nil; bid, _ = cursor.Next() {
		fmt.Printf("%s, key:%d\n", bucketName, bytesToInt64(bid))
		keys[bytesToInt64(bid)] = true
	}
	return keys, nil
}

type TestResourceHolder struct {
}

func (trh *TestResourceHolder) AcquireResources(qm *amqp.QueueMessage) bool {
	return true
}
func (trh *TestResourceHolder) ReleaseResources(qm *amqp.QueueMessage) {

}
