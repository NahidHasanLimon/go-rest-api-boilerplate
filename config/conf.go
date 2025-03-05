package config
import(
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

var AppConfig Config

func LoadConfig(){
	err := godotenv.Load()
	if err!=nil{
		log.Println("⚠️ Warning: No .env file found. Using system environment variables.")
	}

AppConfig = Config{
	DBHost:     getEnv("DB_HOST", "localhost"),
	DBPort:     getEnv("DB_PORT", "5432"),
	DBUser:     getEnv("DB_USER", "postgres"),
	DBPassword: getEnv("DB_PASSWORD", "password"),
	DBName:     getEnv("DB_NAME", "my_database"),
	DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	ServerPort: getEnv("SERVER_PORT", "8070"),
}

}


func getEnv(key, fallback string)string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getBoolEnv(key string, fallback bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if value == "true" {
			return true
		} else if value == "false" {
			return false
		}
		log.Printf("⚠️ Warning: Environment variable %s is set but not a valid boolean. Defaulting to %v.", key, fallback)
	}
	return fallback
}

