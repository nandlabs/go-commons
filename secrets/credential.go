package secrets

type Secret interface {
	Get(key string) string
}
