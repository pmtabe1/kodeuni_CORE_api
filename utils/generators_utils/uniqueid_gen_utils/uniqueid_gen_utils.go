package uniqueid_gen_utils

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/chilts/sid"
	guuid "github.com/google/uuid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
)

func GenShortUUID() string {
	id := shortuuid.New()
	fmt.Printf("github.com/lithammer/shortuuid: %s\n", id)

	return id
}

func GenUUID() string {
	id := guuid.New()
	fmt.Printf("github.com/google/uuid:         %s\n", id.String())

	return id.String()
}

func GenXid() string {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:              %s\n", id.String())

	return id.String()
}

func GenKsuid() string {
	id := ksuid.New()
	fmt.Printf("github.com/segmentio/ksuid:     %s\n", id.String())

	return id.String()
}

func GenBetterGUID() string {
	id := betterguid.New()
	fmt.Printf("github.com/kjk/betterguid:      %s\n", id)

	return id
}

func GenUlid() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("github.com/oklog/ulid:          %s\n", id.String())

	return id.String()
}

func GenSonyflake() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	fmt.Printf("github.com/sony/sonyflake:      %x\n", id)

	return strconv.Itoa(int(id))
}

func GenSid() string {
	id := sid.Id()
	fmt.Printf("github.com/chilts/sid:          %s\n", id)

	return id
}

func GenUUIDv4() string {
	id := uuid.NewV4()
	// if err != nil {
	// 	log.Fatalf("uuid.NewV4() failed with %s\n", err)
	// }
	fmt.Printf("github.com/satori/go.uuid:      %s\n", id)

	return id.String()
}
