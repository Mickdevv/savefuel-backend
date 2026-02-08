package document_categories

import (
	"testing"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/api/auth"
	"github.com/Mickdevv/savefuel-backend/internal/testUtils"
	"github.com/google/uuid"
)

func CreateDocumentCategoryTest(t *testing.T, serverCfg api.ServerConfig, user auth.UserWithTokens) uuid.UUID {

	return uuid.UUID{}
}

func GetDocumentCategoriesTest(t *testing.T, serverCfg api.ServerConfig, user auth.UserWithTokens) uuid.UUID {
	return uuid.UUID{}
}

func UpdateDocumentCategoryTest(t *testing.T, serverCfg api.ServerConfig, user auth.UserWithTokens) uuid.UUID {
	return uuid.UUID{}
}

func DeleteDocumentCategoryTest(t *testing.T, serverCfg api.ServerConfig, user auth.UserWithTokens, id uuid.UUID) {
}

func TestDocumentCategories(t *testing.T) {
	serverCfg := testUtils.TestServerCFG()

	_ = auth.RegisterAndLogin(t, &serverCfg)
}
