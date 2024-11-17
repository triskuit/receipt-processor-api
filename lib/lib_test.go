package lib

import (
	"testing"
)

func TestScoreReceipt(t *testing.T) {

	tests := []struct {
		name string
		r    Receipt
		want int
	}{
		{"example 1",
			Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
				Total: "35.35"},
			28},
		{"example 2",
			Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00"},
			109},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ScoreReceipt(&tt.r)
			if s != tt.want {
				t.Errorf("scoreReceipt() = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestCountAplphaCharacters(t *testing.T) {

	tests := []struct {
		name string
		r    Receipt
		want int
	}{
		{"Target", Receipt{Retailer: "Target"}, 6},
		{"simple string", Receipt{Retailer: "ABC ef"}, 5},
		{"empty string", Receipt{Retailer: "   "}, 0},
		{"non-alpha numeric", Receipt{Retailer: "!&@#"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := 0
			countAlphaCharacters(&tt.r, &s)
			if s != tt.want {
				t.Errorf("countAlphaCharacters() = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestRoundDollarAmount(t *testing.T) {
	tests := []struct {
		name string
		r    Receipt
		want int
	}{
		{"round dollar amount", Receipt{Total: "12.00"}, 50},
		{"square dollar amount", Receipt{Total: "12.01"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := 0
			roundDollarAmount(&tt.r, &s)
			if s != tt.want {
				t.Errorf("roundDollarAmount() = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestMultipleOf25(t *testing.T) {
	tests := []struct {
		name string
		r    Receipt
		want int
	}{
		{"multiple of 0.25", Receipt{Total: "12.75"}, 25},
		{"multiple of 0.25 part deux", Receipt{Total: "12.25"}, 25},
		{"not multiple of 0.25", Receipt{Total: "12.10"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := 0
			multipleOf25(&tt.r, &s)
			if s != tt.want {
				t.Errorf("multipleOf25() = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestTwoItemCount(t *testing.T) {

	tests := []struct {
		name string
		r    Receipt
		want int
	}{
		{"one item", Receipt{Items: []Item{{ShortDescription: "item1", Price: "0.25"}}}, 0},
		{"two items", Receipt{Items: []Item{
			{ShortDescription: "item1", Price: "0.25"},
			{ShortDescription: "item2", Price: "0.50"},
		}}, 5},
		{"four items", Receipt{Items: []Item{
			{ShortDescription: "item1", Price: "0.25"},
			{ShortDescription: "item2", Price: "0.50"},
			{ShortDescription: "item3", Price: "0.75"},
			{ShortDescription: "item4", Price: "1.00"},
		}}, 10},
		{"five items", Receipt{Items: []Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := 0
			twoItemCount(&tt.r, &s)
			if s != tt.want {
				t.Errorf("twoItemCount() = %v, want %v", s, tt.want)

			}
		})
	}
}

func TestDescriptionPoints(t *testing.T) {
	tests := []struct {
		name string
		r    Receipt
		want int
	}{
		{"single item, multiple of 3", Receipt{Items: []Item{
			{ShortDescription: "123456", Price: "10.00"},
		}}, 2},
		{"two items, multiple of 3", Receipt{Items: []Item{
			{ShortDescription: "123", Price: "10.00"},
			{ShortDescription: "123", Price: "10.00"},
		}}, 4},
		{"single item, not a multiple of 3", Receipt{Items: []Item{
			{ShortDescription: "1234", Price: "10.00"},
		}}, 0},
		{"example", Receipt{Items: []Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := 0
			descriptionPoints(&tt.r, &s)
			if s != tt.want {
				t.Errorf("descriptionPoints() = %v, want %v", s, tt.want)

			}
		})
	}
}

func TestOddDate(t *testing.T) {
	tests := []struct {
		name string
		r    Receipt
		want int
	}{
		{"date is odd", Receipt{PurchaseDate: "2022-01-01"}, 6},
		{"date is even", Receipt{PurchaseDate: "2006-02-04"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := 0
			oddDate(&tt.r, &s)
			if s != tt.want {
				t.Errorf("oddDate() = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestEarlyBirdSpecial(t *testing.T) {
	tests := []struct {
		name string
		r    Receipt
		want int
	}{
		{"early bird", Receipt{PurchaseTime: "14:25"}, 10},
		{"very early bird", Receipt{PurchaseTime: "12:25"}, 0},
		{"late bird", Receipt{PurchaseTime: "20:25"}, 0},
		{"on the hour", Receipt{PurchaseTime: "14:00"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := 0
			earlyBirdSpecial(&tt.r, &s)
			if s != tt.want {
				t.Errorf("earlyBirdSpecial() = %v, want %v", s, tt.want)
			}
		})
	}
}
