
package models

import (
    "errors"
    "regexp"
    "time"
)

type Item struct {
    ShortDescription string `json:"shortDescription"`
    Price           string `json:"price"`
}

type Receipt struct {
    Retailer     string `json:"retailer"`
    PurchaseDate string `json:"purchaseDate"`
    PurchaseTime string `json:"purchaseTime"`
    Items        []Item `json:"items"`
    Total        string `json:"total"`
}

type ReceiptResponse struct {
    ID string `json:"id"`
}

type PointsResponse struct {
    Points int `json:"points"`
}

var (
    retailerPattern        = regexp.MustCompile(`^[\w\s\-&]+$`)
    shortDescriptionPattern = regexp.MustCompile(`^[\w\s\-]+$`)
    pricePattern          = regexp.MustCompile(`^\d+\.\d{2}$`)
)

func (r *Receipt) Validate() error {
    if r.Retailer == "" {
        return errors.New("retailer is required. Please verify input.")
    }
    if !retailerPattern.MatchString(r.Retailer) {
        return errors.New("invalid retailer format. Please verify input.")
    }

    if r.PurchaseDate == "" {
        return errors.New("purchase date is required. Please verify input.")
    }
    if _, err := time.Parse("2006-01-02", r.PurchaseDate); err != nil {
        return errors.New("invalid purchase date format, expected YYYY-MM-DD. Please verify input.")
    }

    if r.PurchaseTime == "" {
        return errors.New("purchase time is required. Please verify input.")
    }
    if _, err := time.Parse("15:04", r.PurchaseTime); err != nil {
        return errors.New("invalid purchase time format, expected HH:mm. Please verify input.")
    }

    if len(r.Items) == 0 {
        return errors.New("at least one item is required. Please verify input.")
    }

    for _, item := range r.Items {
        if !shortDescriptionPattern.MatchString(item.ShortDescription) {
            return errors.New("invalid item description format. Please verify input.")
        }
        if !pricePattern.MatchString(item.Price) {
            return errors.New("invalid price format. Please verify input.")
        }
    }

    if r.Total == "" {
        return errors.New("total is required. Please verify input.")
    }
    if !pricePattern.MatchString(r.Total) {
        return errors.New("invalid total format. Please verify input.")
    }

    return nil
}