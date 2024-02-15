package hash_manager

type Config struct {
	Times   uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
	SaltLen int
}
