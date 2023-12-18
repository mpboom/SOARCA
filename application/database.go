package application

import (
	"errors"
	"os"

	mongo "soarca/database/mongodb"
	workflowRepo "soarca/database/workflow"
	routes "soarca/routes"

	"github.com/gin-gonic/gin"
)

func InitialliseDatabase(app *gin.Engine) error {
	mongo.LoadComponent()

	log.Info("SOARCA API Trying to start")
	mongo_uri := os.Getenv("MONGODB_URI")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")

	if mongo_uri == "" || db_username == "" || db_password == "" {
		log.Error("you must set 'MONGODB_URI' or 'DB_USERNAME' or 'DB_PASSWORD' in the environment variable")
		return errors.New("could not obtain required environment settings")
	}
	err := mongo.SetupMongodb(mongo_uri, db_username, db_password)
	if err != nil {
		return err
	}
	// defer database.GetMongoClient().CloseMongoDB()

	workflowRepo := workflowRepo.SetupWorkflowRepository(mongo.GetCacaoRepo(), mongo.DefaultLimitOpts())

	// setup database routes
	err = routes.Database(app, workflowRepo)

	return err
}