package handlers

import (
	"api/cmd/internal/postgresrepo"
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/xuri/excelize/v2"
)

type normalizedData struct {
	Partners      []postgresrepo.CreatePartnerParams      `json:"partners"`
	Customers     []postgresrepo.CreateCustomerParams     `json:"customers"`
	Products      []postgresrepo.CreateProductParams      `json:"products"`
	Skus          []postgresrepo.CreateSkuParams          `json:"skus"`
	Publishers    []postgresrepo.CreatePublisherParams    `json:"publishers"`
	Subscriptions []postgresrepo.CreateSubscriptionParams `json:"subscriptions"`
	Meters        []postgresrepo.CreateMeterParams        `json:"meters"`
	Resources     []postgresrepo.CreateResourceParams     `json:"resources"`
	Entitlement   []postgresrepo.CreateEntitlementParams  `json:"entitlement"`
	Benefits      []postgresrepo.CreateBenefitParams      `json:"benefits"`
	Billings      []postgresrepo.CreateBillingParams      `json:"billings"`
}

type importFileParams struct {
	File *multipart.FileHeader `json:"file"`
}

// ImportFile imports a file and processes it.
// @Summary Import a file
// @Description Import a file xlsx and process it, route protected by JWT
// @Tags import
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to import"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /import [post]
func (h *Handler) ImportFile(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	fileBytes, err := h.getFileBytesFromRequest(r)
	if err != nil {
		h.errorJSON(w, http.StatusInternalServerError, "unable to get file")
		return
	}

	f, err := excelize.OpenReader(bytes.NewReader(fileBytes))
	if err != nil {
		h.errorJSON(w, http.StatusInternalServerError, "Unable to open file")
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.Rows("Planilha1")
	if err != nil {
		h.errorJSON(w, http.StatusInternalServerError, "Unable to get rows")
		return
	}

	normalizedData := h.normalizeRowsData(rows)

	err = h.insertData(normalizedData)

	if err != nil {
		h.errorJSON(w, http.StatusInternalServerError, "Unable to full insert data")
		return
	}

	h.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "File processed successfully",
		"took":    time.Since(startTime).String(),
	})
}

func (h *Handler) insertData(normalizedData normalizedData) error {
	_, err := h.services.CreatePartnersWithCopy(normalizedData.Partners)

	if err != nil {
		log.Println("Error creating partners", err)
		return err
	}

	_, err = h.services.CreateCustomersWithCopy(normalizedData.Customers)

	if err != nil {
		log.Println("Error creating customers", err)
		return err
	}

	_, err = h.services.CreateSkusWithCopy(normalizedData.Skus)

	if err != nil {
		log.Println("Error creating skus", err)
		return err
	}

	_, err = h.services.CreateProductsWithCopy(normalizedData.Products)

	if err != nil {
		log.Println("Error creating products", err)
		return err
	}

	_, err = h.services.CreatePublishersWithCopy(normalizedData.Publishers)

	if err != nil {
		log.Println("Error creating publishers", err)
		return err
	}

	_, err = h.services.CreateSubscriptionsWithCopy(normalizedData.Subscriptions)

	if err != nil {
		log.Println("Error creating subscriptions", err)

		return err
	}

	_, err = h.services.CreateMetersWithCopy(normalizedData.Meters)

	if err != nil {
		log.Println("Error creating meters", err)
		return err
	}

	_, err = h.services.CreateResourcesWithCopy(normalizedData.Resources)

	if err != nil {
		log.Println("Error creating resources", err)
		return err
	}

	_, err = h.services.CreateEntitlementsWithCopy(normalizedData.Entitlement)

	if err != nil {
		log.Println("Error creating entitlements", err)
		return err
	}

	_, err = h.services.CreateBenefitsWithCopy(normalizedData.Benefits)

	if err != nil {
		log.Println("Error creating benefits", err)
		return err
	}

	_, err = h.services.CreateBillingsWithCopy(normalizedData.Billings)

	if err != nil {
		log.Println("Error creating billings", err)
		return err
	}

	return nil
}

func (h *Handler) normalizeRowsData(rows *excelize.Rows) normalizedData {
	var partners []postgresrepo.CreatePartnerParams
	var customers []postgresrepo.CreateCustomerParams
	var products []postgresrepo.CreateProductParams
	var skus []postgresrepo.CreateSkuParams
	var publishers []postgresrepo.CreatePublisherParams
	var subscriptionData []postgresrepo.CreateSubscriptionParams
	var meters []postgresrepo.CreateMeterParams
	var resources []postgresrepo.CreateResourceParams
	var entitlementData []postgresrepo.CreateEntitlementParams
	var benefits []postgresrepo.CreateBenefitParams
	var billingData []postgresrepo.CreateBillingParams
	var normalizedData normalizedData

	seenPartners := make(map[string]bool)
	seenCustomers := make(map[string]bool)
	seenProducts := make(map[string]bool)
	seenSkus := make(map[string]bool)
	seenPublishers := make(map[string]bool)
	seenSubscriptions := make(map[string]bool)
	seenMeter := make(map[string]bool)
	seenResource := make(map[string]bool)
	seenEntitlement := make(map[string]bool)
	seenBenefit := make(map[string]bool)

	for i := 0; rows.Next(); i++ {

		if i == 0 {
			continue
		}

		row, err := rows.Columns()

		if err != nil {
			log.Println("Error getting row columns:", err)
			continue
		}

		processPartner(row, seenPartners, &partners)
		processCustomer(row, seenCustomers, &customers)
		processProduct(row, seenProducts, &products)
		processSku(row, seenSkus, &skus)
		processPublisher(row, seenPublishers, &publishers)
		processSubscription(row, seenSubscriptions, &subscriptionData)
		processMeter(row, seenMeter, &meters)
		processResource(row, seenResource, &resources)
		processEntitlement(row, seenEntitlement, &entitlementData)
		processBenefit(row, seenBenefit, &benefits)
		processBilling(row, &billingData)
	}

	normalizedData.Partners = partners
	normalizedData.Customers = customers
	normalizedData.Products = products
	normalizedData.Skus = skus
	normalizedData.Publishers = publishers
	normalizedData.Subscriptions = subscriptionData
	normalizedData.Meters = meters
	normalizedData.Resources = resources
	normalizedData.Entitlement = entitlementData
	normalizedData.Benefits = benefits
	normalizedData.Billings = billingData

	return normalizedData
}

func processPartner(row []string, seenPartners map[string]bool, partners *[]postgresrepo.CreatePartnerParams) {
	partnerKey := getValue(row, 0)
	if !seenPartners[partnerKey] && partnerKey != "" {
		seenPartners[partnerKey] = true
		*partners = append(*partners, postgresrepo.CreatePartnerParams{
			PartnerKey: getValue(row, 0),
			Name:       getValueAsText(row, 1),
			MpnID:      getValueAsText(row, 6),
		})
	}
}

func processCustomer(row []string, seenCustomers map[string]bool, customers *[]postgresrepo.CreateCustomerParams) {
	customerKey := getValue(row, 2)
	if !seenCustomers[customerKey] && customerKey != "" {
		seenCustomers[customerKey] = true
		*customers = append(*customers, postgresrepo.CreateCustomerParams{

			CustomerKey: getValue(row, 2),
			Name:        getValueAsText(row, 3),
			DomainName:  getValueAsText(row, 4),
			Country:     getValueAsText(row, 5),
			TierToMpnID: getValueAsText(row, 7),
		})
	}
}

func processProduct(row []string, seenProducts map[string]bool, products *[]postgresrepo.CreateProductParams) {
	productKey := getValue(row, 9)
	if !seenProducts[productKey] && productKey != "" {
		seenProducts[productKey] = true
		*products = append(*products, postgresrepo.CreateProductParams{
			ProductKey: getValue(row, 9),
			SkuID:      getValueAsText(row, 10),
			Name:       getValueAsText(row, 13),
		})
	}
}

func processSku(row []string, seenSkus map[string]bool, skus *[]postgresrepo.CreateSkuParams) {
	skuKey := getValue(row, 10)
	if !seenSkus[skuKey] && skuKey != "" {
		seenSkus[skuKey] = true
		*skus = append(*skus, postgresrepo.CreateSkuParams{
			SkuKey:         getValue(row, 10),
			AvailabilityID: getValueAsText(row, 11),
			Name:           getValueAsText(row, 12),
		})
	}
}

func processPublisher(row []string, seenPublishers map[string]bool, publishers *[]postgresrepo.CreatePublisherParams) {
	publisherKey := getValue(row, 15)
	if publisherKey == "" {
		publisherKey = "N/A"
	}

	if !seenPublishers[publisherKey] && publisherKey != "" {
		seenPublishers[publisherKey] = true
		*publishers = append(*publishers, postgresrepo.CreatePublisherParams{
			PublisherKey: publisherKey,
			Name:         getValueAsText(row, 14),
		})
	}
}

func processSubscription(row []string, seenSubscriptions map[string]bool, subscriptions *[]postgresrepo.CreateSubscriptionParams) {
	subscriptionKey := getValue(row, 17)
	if !seenSubscriptions[subscriptionKey] && subscriptionKey != "" {
		seenSubscriptions[subscriptionKey] = true
		*subscriptions = append(*subscriptions, postgresrepo.CreateSubscriptionParams{
			SubscriptionKey: getValue(row, 17),
			Description:     getValueAsText(row, 16),
		})
	}
}

func processMeter(row []string, seenMeter map[string]bool, meters *[]postgresrepo.CreateMeterParams) {
	meterKey := getValue(row, 23)
	if !seenMeter[meterKey] && meterKey != "" {
		seenMeter[meterKey] = true
		*meters = append(*meters, postgresrepo.CreateMeterParams{
			MetersKey:   getValue(row, 23),
			Type:        getValueAsText(row, 21),
			Category:    getValueAsText(row, 22),
			SubCategory: getValueAsText(row, 24),
			Name:        getValueAsText(row, 25),
			Region:      getValueAsText(row, 26),
			Unit:        getValueAsText(row, 27),
		})
	}
}

func processResource(row []string, seenResource map[string]bool, resources *[]postgresrepo.CreateResourceParams) {
	resourceKey := getValue(row, 31)
	if !seenResource[resourceKey] && resourceKey != "" {
		seenResource[resourceKey] = true
		*resources = append(*resources, postgresrepo.CreateResourceParams{
			Uri:             getValue(row, 31),
			Location:        getValueAsText(row, 28),
			ConsumedService: getValueAsText(row, 29),
			ResourceGroup:   getValueAsText(row, 30),
			Info1:           getValueAsText(row, 40),
			Info2:           getValueAsText(row, 41),
			Tags:            getValueAsText(row, 42),
			AdditionalInfo:  getValueAsText(row, 43),
		})
	}
}

func processEntitlement(row []string, seenEntitlement map[string]bool, entitlements *[]postgresrepo.CreateEntitlementParams) {
	entitlementKey := getValue(row, 47)
	if !seenEntitlement[entitlementKey] && entitlementKey != "" {
		seenEntitlement[entitlementKey] = true
		*entitlements = append(*entitlements, postgresrepo.CreateEntitlementParams{
			EntitlementKey: getValue(row, 47),
			Description:    getValueAsText(row, 48),
		})
	}
}

func processBenefit(row []string, seenBenefit map[string]bool, benefits *[]postgresrepo.CreateBenefitParams) {
	if len(row) < 55 {
		return
	}

	benefitKey := getValue(row, 53)

	if benefitKey == "" {
		benefitKey = "N/A"
	}

	if !seenBenefit[benefitKey] && benefitKey != "" {
		seenBenefit[benefitKey] = true
		*benefits = append(*benefits, postgresrepo.CreateBenefitParams{
			BenefitKey:     benefitKey,
			BenefitOrderID: getValueAsText(row, 52),
			BenefitType:    getValueAsText(row, 54),
		})
	}
}

func processBilling(row []string, billings *[]postgresrepo.CreateBillingParams) {
	publisherKey := getValueAsText(row, 15)
	benefitKey := getValueAsText(row, 53)

	if publisherKey.String == "" {
		publisherKey.String = "N/A"
	}

	if benefitKey.String == "" {
		benefitKey.String = "N/A"
	}

	*billings = append(*billings, postgresrepo.CreateBillingParams{
		PartnerKey:             getValueAsText(row, 0),
		CustomerKey:            getValueAsText(row, 2),
		ProductKey:             getValueAsText(row, 9),
		PublisherKey:           publisherKey,
		SubscriptionKey:        getValueAsText(row, 17),
		MetersKey:              getValueAsText(row, 23),
		ResourceUri:            getValueAsText(row, 31),
		EntitlementKey:         getValueAsText(row, 47),
		BenefitKey:             benefitKey,
		Invoice:                getValueAsText(row, 8),
		UnitPrice:              getValueAsNumeric(row, 33),
		Quantity:               getValueAsNumeric(row, 34),
		UnitType:               getValueAsText(row, 35),
		BillingPreTaxTotal:     getValueAsNumeric(row, 36),
		BillingCurrency:        getValueAsText(row, 37),
		PricingPreTaxTotal:     getValueAsNumeric(row, 38),
		PricingCurrency:        getValueAsText(row, 39),
		EffectiveUnitPrice:     getValueAsNumeric(row, 44),
		PcToBcExchangeRate:     getValueAsNumeric(row, 45),
		PcToBcExchangeRateDate: getValueAsDate(row, 46),
		ChargeStartDate:        getValueAsDate(row, 18),
		ChargeEndDate:          getValueAsDate(row, 19),
		UsageDate:              getValueMondayDayYearAsDate(row, 20),
		ChargeType:             getValueAsText(row, 32),
	})
}

func getValue(row []string, index int) string {
	if index < len(row) {
		return row[index]
	}
	return ""
}

func getValueAsText(row []string, index int) pgtype.Text {
	if index < len(row) {
		var text pgtype.Text

		if err := text.Scan(row[index]); err != nil {
			log.Println(err, text)
			log.Println(">>>>>>>>", index, "<<<<<<<<")
			log.Println(">>>>>>>>", row[index], "<<<<<<<<")
			panic(err)
		}
		return text
	}
	return pgtype.Text{}
}

func getValueAsNumeric(row []string, index int) pgtype.Numeric {
	if index < len(row) {
		var numeric pgtype.Numeric
		value := strings.TrimSpace(row[index])

		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println(err, numeric)
			return pgtype.Numeric{}
		}

		if err := numeric.Scan(strconv.FormatFloat(floatValue, 'f', -1, 64)); err != nil {
			log.Println(err, numeric)
			return pgtype.Numeric{}
		}

		return numeric
	}
	return pgtype.Numeric{}
}

func getValueAsDate(row []string, index int) pgtype.Date {
	if index < len(row) {
		var date pgtype.Date
		value := strings.TrimSpace(row[index])

		parsedDate, err := time.Parse("01-02-06", value)

		if err != nil {
			log.Println("Error parsing date", err)
			return pgtype.Date{}
		}

		if err := date.Scan(parsedDate); err != nil {
			log.Println("Error parsing date", err)
			return pgtype.Date{}
		}
		return date
	}
	return pgtype.Date{}
}

func getValueMondayDayYearAsDate(row []string, index int) pgtype.Date {
	if index < len(row) {
		var date pgtype.Date
		value := strings.TrimSpace(row[index])

		layouts := []string{
			"01/02/2006", // MM/DD/YYYY
			"02-01-06",   // DD-MM-YY
		}

		var parsedDate time.Time
		var err error

		for _, layout := range layouts {
			parsedDate, err = time.Parse(layout, value)
			if err == nil {
				break
			}
		}

		if err != nil {
			log.Println(err)
			return pgtype.Date{}
		}

		if err := date.Scan(parsedDate); err != nil {
			log.Println(err)
			return pgtype.Date{}
		}
		return date
	}
	return pgtype.Date{}
}

func (h *Handler) getFileBytesFromRequest(r *http.Request) ([]byte, error) {
	err := r.ParseMultipartForm(100 << 20) // 100 MB
	if err != nil {
		return nil, fmt.Errorf("Unable to parse multipart form: %v", err)
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("Unable to get file: %v", err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to read file: %v", err)
	}

	return fileBytes, nil
}
