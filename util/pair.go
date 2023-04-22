package util

type Pair[K comparable, V any] struct {
	Key   K `json:"key" bson:"_id"`
	Value V `json:"value" bson:"value"`
}

type Pairs[K comparable, V any] []*Pair[K, V]

func (pairs Pairs[K, V]) ToMap() map[K]V {
	if pairs == nil {
		return nil
	}

	var m = make(map[K]V)
	for _, pair := range pairs {
		m[pair.Key] = pair.Value
	}
	return m
}
