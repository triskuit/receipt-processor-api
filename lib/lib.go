package lib

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	RoundDollarPoints     = 50
	MultipleOf25Points    = 25
	PairOfItemsPoints     = 5
	OddDayPoints          = 6
	EarlyBirdPoints       = 10
	DescriptionMultiplier = 0.2
)

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required" validate:"regexp=^[\\w\\s\\-&]+$"`
	PurchaseDate string `json:"purchaseDate" binding:"required" validate:"regexp=^\\d{4}-\\d{2}-\\d{2}$"`
	PurchaseTime string `json:"purchaseTime" binding:"required" validate:"regexp=^([01]\\d|2[0-3]):[0-5]\\d$"`
	Total        string `json:"total" binding:"required" validate:"regexp=^\\d+\\.\\d{2}$"`
	Items        []Item `json:"items" binding:"required"`
}

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required" validate:"regexp=^[\\w\\s\\-]+$"`
	Price            string `json:"price" binding:"required" validate:"regexp=^\\d+\\.\\d{2}$"`
}

// Returns a score for a given receipt
func ScoreReceipt(r *Receipt) (score int) {
	score = 0
	countAlphaCharacters(r, &score)
	roundDollarAmount(r, &score)
	multipleOf25(r, &score)
	twoItemCount(r, &score)
	descriptionPoints(r, &score)
	oddDate(r, &score)
	earlyBirdSpecial(r, &score)
	return
}

// RULES

// Updates the score by adding 1 point for every alphanumeric character in the retailer name
func countAlphaCharacters(r *Receipt, s *int) {
	count := 0
	for _, char := range r.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	*s = *s + count
}

// Updates the score by adding 50 points if the total is a round dollar amount, zero otherwise
func roundDollarAmount(r *Receipt, s *int) {
	if strings.HasSuffix(r.Total, ".00") {
		*s = *s + RoundDollarPoints
	}
}

// Updates the score by adding 25 points if the total is a multiple of 0.25
func multipleOf25(r *Receipt, s *int) {
	float, _ := strconv.ParseFloat(r.Total, 32)
	if math.Mod(float, 0.25) == 0 {
		*s = *s + MultipleOf25Points
	}
}

// Updates the score by adding 5 points for every two items on the receipt
func twoItemCount(r *Receipt, s *int) {
	count := int(float64(len(r.Items) / 2))
	*s = *s + (PairOfItemsPoints * count)
}

// Updates the score by adding X points where X is the price multiplied by 0.2
// and rounded up IF the trimmed length of the item description is a
// multiple of 3 PER ITEM
func descriptionPoints(r *Receipt, s *int) {
	for _, item := range r.Items {
		strLength := len(strings.Trim(item.ShortDescription, " "))
		if strLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 32)
			points := int(math.Ceil(price * DescriptionMultiplier))
			*s = *s + points
		}
	}
}

// Updates the score by adding 6 points if the day in the purchase date is odd
func oddDate(r *Receipt, s *int) {
	date, _ := time.Parse("2006-01-02", r.PurchaseDate)
	if date.Day()%2 != 0 {
		*s = *s + OddDayPoints
	}
}

// Updates the score by adding 10 points if the time of purchase is after 2:00pm and before 4:00pm
func earlyBirdSpecial(r *Receipt, s *int) {
	t, _ := time.Parse("15:04", r.PurchaseTime)
	if t.Hour() >= 14 && t.Hour() < 16 {
		*s = *s + EarlyBirdPoints
	}
}
