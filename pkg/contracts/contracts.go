package contracts

/**
 * Topics.
 */
const (
	LogsTopic           = "logs"
	ColorTopic          = "color"
	FamilyTopic         = "family"
	ProductTopic        = "product"
	CollectionTopic     = "collection"
	ProductVariantTopic = "product-variant"
)

/**
 * Log.
 */
type LogMessage struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

/**
 * Color.
 */
type ColorMessage struct {
	ID   string `json:"id"`
	Hex  string `json:"hex"`
	Name string `json:"name"`
}

/**
 * Family.
 */
type FamilyMessage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/**
 * Collection.
 */
type CollectionMessage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/**
 * Product.
 */
type ProductMessage struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/**
 * Product variant.
 */
type ProductVariantMessage struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ProductID string  `json:"product_id"`
}
