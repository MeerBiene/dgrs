package dgrc

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
)

// Options defines State preferences.
type Options struct {
	// You can pass a pre-initialized redis instance
	// if you already have one.
	RedisClient *redis.Client

	// Redis client options to connect to a redis
	// instance.
	RedisOptions redis.Options

	// Fetch requested values directly from the Discord API
	// and store them in the cache.
	FetchAndStore bool

	// You can specify a timeout period for redis commands.
	// If not set, no timeout will be used.
	CommandTimeout time.Duration
}

type State struct {
	options *Options
	client  *redis.Client
	session *discordgo.Session
}

func New(session *discordgo.Session, opts Options) (s *State) {
	s = &State{}

	s.session = session
	if opts.RedisClient != nil {
		s.client = opts.RedisClient
	} else {
		s.client = redis.NewClient(&opts.RedisOptions)
	}

	s.options = &opts

	return
}
