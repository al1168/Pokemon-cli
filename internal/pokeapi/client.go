package pokeapi
import (
	"net/http"
	"time"
	"github.com/al1168/Pokemon-cli/internal/pokecache"
)
type Client struct{
	client http.Client
	cache pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client{
	return Client{
		client : http.Client{
			Timeout: timeout,
		},
		cache : pokecache.NewCache(cacheInterval),
	}
}