package kvs

var KVS map[string]interface{}

func InitKVS() map[string]interface{} {
	KVS = make(map[string]interface{})
	return KVS
}
