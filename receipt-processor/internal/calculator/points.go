package calculator

import (
    "math"
    "regexp"
    "strconv"
    "strings"
    "time"
    "receipt-processor/api/models"
)

func CalculatePoints(receipt *models.Receipt) (int, error) {
    points := 0
    alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
    matches := alphanumeric.FindAllString(receipt.Retailer, -1)
    points += len(matches)


    total, err := strconv.ParseFloat(receipt.Total, 64)
    if err != nil {
        return 0, err
    }
    if total == float64(int(total)) {
        points += 50
    }

    if math.Mod(total*100, 25) == 0 {
        points += 25
    }

    points += (len(receipt.Items) / 2) * 5


    for _, item := range receipt.Items {
        trimmedLen := len(strings.TrimSpace(item.ShortDescription))
        if trimmedLen%3 == 0 {
            price, err := strconv.ParseFloat(item.Price, 64)
            if err != nil {
                return 0, err
            }
            points += int(math.Ceil(price * 0.2))
        }
    }


    if total > 10.0 {
        points += 5
    }

    purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
    if err != nil {
        return 0, err
    }
    if purchaseDate.Day()%2 == 1 {
        points += 6
    }

    purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
    if err != nil {
        return 0, err
    }
    if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
        points += 10
    }

    return points, nil
}
