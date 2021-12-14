package app_databases

import (
	"os"
	"testing"

	"github.com/paulmsegeya/subscription/config/app_config"
	"github.com/stretchr/testify/require"
)

func TestConnectFromDB(t *testing.T) {
	got := New().ConnectFromDB()
	require.NotNilf(t, got, "Expected non nil but got %v instead", got)

}

func TestConnectToDB(t *testing.T) {
	got := New().ConnectToDB()
	require.Nilf(t, got, "Expected non nil but got %v instead", got)

}

func TestDBConnection(t *testing.T) {
	got := New().DBConnection()
	require.Nilf(t, got, "Expected non nil but got %v instead", got)

}

func TestConnectionStringFrom(t *testing.T) {

	conf := app_config.New()
	got := New().ConnectionStringFrom(conf.ReadConfiguration())
	require.NotEmptyf(t, got, "Expected non nil but got %v instead", got)

}

func TestPingDatabaseStatus(t *testing.T) {
	os.Setenv("POS_CONFIG", "./config.%v.json")
	instanceConn := New().ConnectToDB()
	got := New().PingDatabaseStatus(instanceConn)
	require.Truef(t, got, "Expected true but got %v instead", got)

}
