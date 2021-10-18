package vars

var (
	ConfigDir     string
	DefaultConfig = &Config{}
)

type Config struct {
	FirebaseProjectID     string `mapstructure:"firebase_project_id"`
	FirestoreCollectionID string `mapstructure:"firestore_collection_id"`
	OCBCAPIClientID       string `mapstructure:"ocbcapi_client_id"`
	CallbackHost          string `mapstructure:"callback_host"`
}
