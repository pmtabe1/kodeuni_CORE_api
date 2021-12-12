package timeutils

import (
	"fmt"
	"log"
	"testing"
	"time"

	//"time"

 	"github.com/stretchr/testify/require"
)

func TestFromJavaLongMilliseconds(t *testing.T) {
	got := FromJavaLongMilliseconds(1622926800000)
	log.Println("Time From Long " + fmt.Sprintf("%v", got))
	require.Nilf(t, got, "Expected Non Nil but got %v  instead", got)

}

func TestFromJavaLongMillisecondsGetDate(t *testing.T) {
	got := FromJavaLongMillisecondsGetDate(1633204092000)
	log.Println("Time From Long " + fmt.Sprintf("%v", got))
	require.NotEmptyf(t, got, "Expected Non Nil but got %v  instead", got)
}

func TestFromJavaLongMillisecondsGetDateTime(t *testing.T) {
	got := FromJavaLongMillisecondsGetDate(1622926800000)
	log.Println("Time From Long " + fmt.Sprintf("%v", got))
	require.NotEmptyf(t, got, "Expected Non Nil but got %v  instead", got)
}

func TestFromJavaLongMillisecondsGetTime(t *testing.T) {
	got := FromJavaLongMillisecondsGetTime(time.Now().Unix())
	log.Println("Time From Long " + fmt.Sprintf("%v", got))
	require.NotEmptyf(t, got, "Expected Non Nil but got %v  instead", got)



}


func TestGetTokenRemainingValidity(t *testing.T)  {
	var claims map[string]int64
    claims = make(map[string]int64)
    claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	
    //x := getTokenRemainingValidity(claims["exp"])
	got:=GetTokenRemainingValidity(claims["exp"])
	require.Greaterf(t,got,8000, "Expected to be Greater but got %v  instead", got)

}

func TestGetFineractDate(t *testing.T)  {
	
	got:=GetFineractDate()
	require.Nilf(t,got,"Expected  got %v",got)
}


func TestGetTRATokenRemainingValidity(t *testing.T)  {
	// tin:="112318380"
	// registration := registration_repository.New().GetByTin(tin).Registration
	// got:=GetTRATokenRemainingValidity(registration.ExpiresIn,tin)
	// require.NotEmptyf(t, got, "Expected Non Nil but got %v  instead", got)
	// require.NotNil(t, got, "Expected Non Nil but got %v  instead", got)
	// require.Greater(t,got,registration.ExpiresIn, "Expected to be Greater but got %v  instead", got)



	
}
