package rentals

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/apkriege/outdoorsy-backend/common/db"
	"github.com/apkriege/outdoorsy-backend/common/models"
	"github.com/apkriege/outdoorsy-backend/common/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var sortByMap = map[string]string{
	"price": "price_per_day",
}

type Rental struct {
	ID              int            `json:"id"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Type            string         `json:"type"`
	Make            string         `json:"make"`
	Model           string         `json:"model"`
	Year            int            `json:"year"`
	Length          float64        `json:"length"`
	Sleeps          int            `json:"sleeps"`
	PrimaryImageURL string         `json:"primary_image_url"`
	Price           RentalPrice    `json:"price"`
	Location        RentalLocation `json:"location"`
	User            models.User    `json:"user"`
}

type RentalPrice struct {
	Day int `json:"day"`
}

type RentalLocation struct {
	City    string  `json:"city"`
	State   string  `json:"state"`
	Zip     string  `json:"zip"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type RentalQuery struct {
	Query *gorm.DB
}

// GetRental endpoint returns a single rental
func GetRental(c *gin.Context) {
	db := db.GetDB()
	id := c.Param("id")

	var rental models.Rental
	db.Table("rentals").Where("id = ?", id).Preload("User").Find(&rental)

	r := modelRental(&rental)
	c.JSON(http.StatusOK, r)
}

// GetRentals endpoint returns a list of rentals
func GetRentals(c *gin.Context) {
	db := db.GetDB()
	queryString := c.Request.URL.Query()

	query := NewRentalQuery(db)
	applyFilters(query, queryString)

	rentals, err := query.FindRentals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching rentals"})
		return
	}

	if len(rentals) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No rentals found"})
		return
	}

	modeledRentals := []Rental{}
	for _, r := range rentals {
		modeledRentals = append(modeledRentals, modelRental(&r))
	}

	utils.ReturnSuccess(c, modeledRentals)
}

// NewRentalQuery returns a new RentalQuery struct
func NewRentalQuery(db *gorm.DB) *RentalQuery {
	return &RentalQuery{
		Query: db.Table("rentals").Preload("User"),
	}
}

// applyFilters applies the filters to the query
func applyFilters(query *RentalQuery, queryString map[string][]string) {
	for key, values := range queryString {
		value := values[0]
		if value != "" {
			switch key {
			case "near":
				query.addNearFilter(value)
			case "ids":
				query.addIdsFilter(value)
			case "price_min":
				query.addPriceMinFilter(value)
			case "price_max":
				query.addPriceMaxFilter(value)
			case "limit":
				query.addLimitFilter(value)
			case "offset":
				query.addOffsetFilter(value)
			case "sort":
				query.addSortFilter(value)
			}
		}
	}
}

// FindRentals returns a list of rentals
func (q *RentalQuery) FindRentals() ([]models.Rental, error) {
	var rentals []models.Rental
	err := q.Query.Find(&rentals).Error
	if err != nil {
		return nil, err
	}
	return rentals, nil
}

// DATA FILTERS
func (q *RentalQuery) addNearFilter(near string) {
	coords := strings.Split(near, ",")
	latitude := coords[0]
	longitude := coords[1]

	nearString := fmt.Sprintf(`
		ST_DWithin(
			ST_MakePoint(lat, lng)::geography,
			ST_MakePoint(%s, %s)::geography,
			%d * 1609.34
		)`, latitude, longitude, 100)

	q.Query.Where(nearString)
}

func (q *RentalQuery) addPriceMinFilter(priceMin string) {
	priceMinInt, _ := strconv.Atoi(priceMin)
	q.Query.Where("price_per_day >= ?", priceMinInt)
}

func (q *RentalQuery) addPriceMaxFilter(priceMax string) {
	priceMaxInt, _ := strconv.Atoi(priceMax)
	q.Query.Where("price_per_day <= ?", priceMaxInt)
}

func (q *RentalQuery) addLimitFilter(limit string) {
	limitInt, _ := strconv.Atoi(limit)
	q.Query.Limit(limitInt)
}

func (q *RentalQuery) addOffsetFilter(offset string) {
	offsetInt, _ := strconv.Atoi(offset)
	q.Query.Offset(offsetInt)
}

func (q *RentalQuery) addSortFilter(sort string) {
	sortBy := sortByMap[sort]
	q.Query.Order(sortBy)
}

func (q *RentalQuery) addIdsFilter(ids string) {
	idsArr := strings.Split(ids, ",")
	q.Query.Where("id IN (?)", idsArr)
}

func modelRental(r *models.Rental) Rental {
	rental := Rental{
		ID:              r.ID,
		Name:            r.Name,
		Description:     r.Description,
		Type:            r.Type,
		Make:            r.VehicleMake,
		Model:           r.VehicleModel,
		Year:            r.VehicleYear,
		Length:          r.VehicleLength,
		Sleeps:          r.Sleeps,
		PrimaryImageURL: r.PrimaryImageUrl,
	}

	price := RentalPrice{
		Day: r.PricePerDay,
	}

	location := RentalLocation{
		City:    r.HomeCity,
		State:   r.HomeState,
		Zip:     r.HomeZip,
		Country: r.HomeCountry,
		Lat:     r.Lat,
		Lng:     r.Lng,
	}

	user := models.User{
		ID:        r.UserId,
		FirstName: r.User.FirstName,
		LastName:  r.User.LastName,
	}

	rental.Price = price
	rental.Location = location
	rental.User = user

	return rental
}
