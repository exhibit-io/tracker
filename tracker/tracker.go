package tracker

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/julienschmidt/httprouter"

	// Import the config package
	"github.com/exhibit-io/tracker/config"
)

var ctx = context.Background()
var rdb *redis.Client
var hash = sha256.New()

func Init(config *config.Config) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.GetAddr(), // Redis server address
		Password: config.Redis.Password,
	})
	if rdb.Ping(ctx).Err() != nil {
		log.Fatal("Failed to connect to Redis")
		panic("Failed to connect to Redis")
	}
	log.Println("Connected to Redis on " + config.Redis.GetAddr())

}

func TrackerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// get the path from the query parameter
	path := r.URL.Query().Get("path")

	// Get IP address from headers or remote address
	ip := getIPAddress(r)

	// Check if the fingerprint cookie exists
	cookie, err := r.Cookie("fingerprint")
	fingerprint := ""
	visits := 1

	if err == nil {
		// If the cookie exists, use the fingerprint from the cookie
		fingerprint = cookie.Value

		// Increment the number of visits for the path
		visits = int(rdb.IncrBy(ctx, getRedisKey(fingerprint, path), 1).Val())

	} else {
		// Create a new fingerprint using IP and User-Agent if cookie doesn't exist
		fingerprint = createFingerprint(ip, r.UserAgent())

		// Store a map on redis with the fingerprint as the key and the path as the value and the number of times the path was visited
		rdb.Set(ctx, getRedisKey(fingerprint, path), visits, 0)

		// Set the fingerprint as a cookie
		http.SetCookie(w, &http.Cookie{
			Name:  "fingerprint",
			Value: fingerprint,
			Path:  "/",
			Domain: "onrender.com",
			// Secure: true, // Uncomment if using HTTPS
			// HttpOnly: true, // Uncomment to prevent client-side scripts from accessing the cookie
		})
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"ip":          ip,
		"userAgent":   r.UserAgent(),
		"fingerprint": fingerprint,
		"path":        path,
		"visits":      fmt.Sprintf("%d", visits),
	}
	json.NewEncoder(w).Encode(response)
	log.Printf(">> %-7s %03d /%s", fingerprint[:7], visits, path)
}

func getRedisKey(fingerprint, path string) string {
	return fmt.Sprintf("tracker:%s:%s", fingerprint, path)
}

func getIPAddress(r *http.Request) string {
	// Check X-Forwarded-For header first
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		ips := strings.Split(forwardedFor, ",")
		if len(ips) > 0 && ips[0] != "" {
			return strings.TrimSpace(ips[0])
		}
	}

	// Fall back to X-Real-IP header
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}

	// Default to RemoteAddr
	return r.RemoteAddr
}

func createFingerprint(ip string, userAgent string) string {
	hash.Write([]byte(ip + userAgent))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
