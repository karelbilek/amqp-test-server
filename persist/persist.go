package persist

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	bolt "go.etcd.io/bbolt"
)

type UnmarshalerFactory interface {
	New() proto.Unmarshaler
}

//
//            Persist
//

func PersistOne(db *bolt.DB, bucketName []byte, key string, obj proto.Marshaler) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil { // pragma: nocover
			return fmt.Errorf("create bucket: %s", err)
		}
		return PersistOneBoltTx(bucket, key, obj)
	})
}

func PersistOneBoltTx(bucket *bolt.Bucket, key string, obj proto.Marshaler) error {
	exBytes, err := obj.Marshal()
	if err != nil { // pragma: nocover -- no idea how to produce this error
		return fmt.Errorf("Could not marshal object")
	}
	return bucket.Put([]byte(key), exBytes)
}

func PersistMany(db *bolt.DB, bucketName []byte, objs map[string]proto.Marshaler) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil { // pragma: nocover
			return fmt.Errorf("create bucket: %s", err)
		}
		return PersistManyBoltTx(bucket, objs)
	})
}

func PersistManyBoltTx(bucket *bolt.Bucket, objs map[string]proto.Marshaler) error {
	for key, obj := range objs {
		err := PersistOneBoltTx(bucket, key, obj)
		if err != nil {
			return err
		}
	}
	return nil
}

//
//                    Load
//

func LoadOne(db *bolt.DB, bucketName []byte, key string, obj proto.Unmarshaler) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("Bucket not found: '%s'", bucket)
		}
		return LoadOneBoltTx(bucket, key, obj)
	})
}

func LoadOneBoltTx(bucket *bolt.Bucket, key string, obj proto.Unmarshaler) error {
	objBytes := bucket.Get([]byte(key))
	if objBytes == nil {
		return fmt.Errorf("Key not found: '%s'", key)
	}
	err := obj.Unmarshal(objBytes)
	if err != nil {
		return fmt.Errorf("Could not unmarshal key %s", key)
	}
	return nil
}

func LoadMany(db *bolt.DB, bucketName []byte, objs map[string]proto.Unmarshaler) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil { // pragma: nocover
			return fmt.Errorf("create bucket: '%s'", bucket)
		}
		return LoadManyBoltTx(bucket, objs)
	})
}

func LoadManyBoltTx(bucket *bolt.Bucket, objs map[string]proto.Unmarshaler) error {
	for key, obj := range objs {
		err := LoadOneBoltTx(bucket, key, obj)
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadAll(db *bolt.DB, bucket []byte, factory UnmarshalerFactory) (map[string]proto.Unmarshaler, error) {
	ret := make(map[string]proto.Unmarshaler)
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket)
		if bucket == nil {
			return nil
		}
		// iterate through queues
		cursor := bucket.Cursor()
		for name, data := cursor.First(); name != nil; name, data = cursor.Next() {
			obj := factory.New()
			err := obj.Unmarshal(data)
			if err != nil {
				return fmt.Errorf("Could not unmarshal key %s", string(name))
			}
			ret[string(name)] = obj
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//
//                      Depersist
//

func DepersistOne(db *bolt.DB, bucketName []byte, key string) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("Bucket not found: '%s'", bucket)
		}
		return DepersistOneBoltTx(bucket, key)
	})
}

func DepersistOneBoltTx(bucket *bolt.Bucket, key string) error {
	return bucket.Delete([]byte(key))
}

func DepersistMany(db *bolt.DB, bucketName []byte, keys map[string]bool) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil { // pragma: nocover
			return fmt.Errorf("create bucket: '%s'", bucket)
		}
		return DepersistManyBoltTx(bucket, keys)
	})
}

func DepersistManyBoltTx(bucket *bolt.Bucket, keys map[string]bool) error {
	for key, _ := range keys {
		err := DepersistOneBoltTx(bucket, key)
		if err != nil {
			return err
		}
	}
	return nil
}
