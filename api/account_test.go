package api

import (
	"bytes"
	"fmt"
	mockdb "learn/banking/db/mock"
	db "learn/banking/db/sqlc"
	"learn/banking/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	server := NewServer(store)
	recoder := httptest.NewRecorder()

	url := fmt.Sprintf("/account/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recoder, request)
	require.Equal(t, http.StatusOK, recoder.Code)

}

func randomAccount() db.Account {
	return db.Account{
		ID:       utils.RandomInt(1, 1000),
		Owner:    utils.RandomOwner(),
		Currency: utils.RandomCurrency(),
		Balance:  utils.RandomMoney(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {

}
