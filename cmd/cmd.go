package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	"monolith/internal/usecase/hash_manager"
	"monolith/internal/usecase/token_manager"
)

func init() {
	mustInitEnv()
}

func mustInitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func MustInitPostgresql() *sqlx.DB {
	conn, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASS"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DBNAME"),
			os.Getenv("POSTGRES_SSLMODE"),
		),
	)
	if err != nil {
		panic(err)
	}
	if err = conn.Ping(); err != nil {
		panic(err)
	}

	return conn
}

func MustInitPostgresqlRO() *sqlx.DB {
	conn, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASS"),
			os.Getenv("POSTGRES_HOST_RO"),
			os.Getenv("POSTGRES_PORT_RO"),
			os.Getenv("POSTGRES_DBNAME"),
			os.Getenv("POSTGRES_SSLMODE"),
		),
	)
	if err != nil {
		panic(err)
	}
	if err = conn.Ping(); err != nil {
		panic(err)
	}

	return conn
}

func MustInitHasherConfig() hash_manager.Config {
	var conf hash_manager.Config

	if times, err := strconv.ParseUint(os.Getenv("HASHER_TIMES"), 10, 32); err == nil {
		conf.Times = uint32(times)
	} else {
		panic(err)
	}

	if memory, err := strconv.ParseUint(os.Getenv("HASHER_MEMORY"), 10, 32); err == nil {
		conf.Memory = uint32(memory)
	} else {
		panic(err)
	}

	if threads, err := strconv.ParseUint(os.Getenv("HASHER_THREADS"), 10, 8); err == nil {
		conf.Threads = uint8(threads)
	} else {
		panic(err)
	}

	if keyLen, err := strconv.ParseUint(os.Getenv("HASHER_KEY_LEN"), 10, 32); err == nil {
		conf.KeyLen = uint32(keyLen)
	} else {
		panic(err)
	}

	if saltLen, err := strconv.ParseInt(os.Getenv("HASHER_SALT_LEN"), 10, 64); err == nil {
		conf.SaltLen = int(saltLen)
	} else {
		panic(err)
	}

	return conf
}

func MustInitTokenManagerConfig() token_manager.Config {
	var privateKey, publicKey string
	var ok bool

	if privateKey, ok = os.LookupEnv("TOKEN_PRIVATE_KEY"); !ok {
		panic("TOKEN_PRIVATE_KEY not set")
	}

	if publicKey, ok = os.LookupEnv("TOKEN_PUBLIC_KEY"); !ok {
		panic("TOKEN_PUBLIC_KEY not set")
	}

	return token_manager.Config{
		TtlAccessToken: token_manager.TtlAccessTokenDefault,
		PrivateKey:     privateKey,
		PublicKey:      publicKey,
	}
}

func MustInitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
}
