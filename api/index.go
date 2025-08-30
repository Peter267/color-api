package handler

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"  // Register GIF decoder
	_ "image/jpeg" // Register JPEG decoder
	_ "image/png"  // Register PNG decoder
	"net/http"

	"github.com/nfnt/resize"
	_ "golang.org/x/image/webp" // Register WebP decoder
)

// Response struct for our JSON output
type Response struct {
	RGB string `json:"RGB"`
}

// writeError is a helper to create a standard JSON error response.
func writeError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// Handler is the main entry point for the Vercel serverless function.
func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. Get the image URL from the query parameters
	imageURL := r.URL.Query().Get("url")
	if imageURL == "" {
		writeError(w, "Missing image url", http.StatusBadRequest)
		return
	}

	// 2. Download the image
	resp, err := http.Get(imageURL)
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to download image: %v", err), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		writeError(w, fmt.Sprintf("Image source returned status: %s", resp.Status), http.StatusBadRequest)
		return
	}

	// 3. Decode the image. The blank imports above handle format detection.
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to decode image: %v", err), http.StatusInternalServerError)
		return
	}

	// 4. --- PERFORMANCE OPTIMIZATION ---
	// Resize the image to a small thumbnail for faster processing.
	// A width of 100px is more than enough to find the dominant color.
	thumbnail := resize.Resize(100, 0, img, resize.Lanczos3)

	// 5. Find the dominant color by creating a color histogram
	colorCounts := make(map[color.Color]int)
	bounds := thumbnail.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			colorCounts[thumbnail.At(x, y)]++
		}
	}

	if len(colorCounts) == 0 {
		writeError(w, "Could not extract any colors from the image", http.StatusInternalServerError)
		return
	}

	// Find the color with the highest count
	var dominantColor color.Color
	maxCount := 0
	for c, count := range colorCounts {
		if count > maxCount {
			maxCount = count
			dominantColor = c
		}
	}

	// 6. Convert the color to a HEX string
	// The RGBA() method returns 16-bit values, so we shift right by 8 to get 8-bit values.
	rVal, gVal, bVal, _ := dominantColor.RGBA()
	hexColor := fmt.Sprintf("#%02x%02x%02x", rVal>>8, gVal>>8, bVal>>8)

	// 7. Send the successful JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{RGB: hexColor})
}