package firebase_client

// import (
// 	"context"
// 	"os"

// 	firebase "firebase.google.com/go"
// 	"firebase.google.com/go/messaging"
// 	"google.golang.org/api/option"
// )
  

// type IFirebaseClient interface {
	
// }

// type FirebaseClient struct {
	
// }


// func New()  *FirebaseClient {
	

// 	  // There are different ways to add credentials on init.
// 	  // if we have a path to the JSON credentials file, we use the GOOGLE_APPLICATION_CREDENTIALS env var
// 	  os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", c.Firebase.Credentials)
// 	  // or pass the file path directly
// 	  opts := []option.ClientOption{option.WithCredentialsFile("creds.json")}
	  
// 	  // if we have a raw JSON credentials value, we use the FIREBASE_CONFIG env var
// 	  os.Setenv("FIREBASE_CONFIG", "{...}")
	  
// 	  // or we can pass the raw JSON value directly as an option
// 	  opts := []option.ClientOption{option.WithCredentialsJSON([]byte("{...}"))}
	  
	  
// 	  app, err := firebase.NewApp(ctx, nil, opts...)
// 	  if err != nil {
// 		  log.Fatalf("new firebase app: %s", err)
// 	  }
	  
// 	  fcmClient, err := app.Messaging(context.TODO())
// 	  if err != nil {
// 		  log.Fatalf("messaging: %s", err) 
// 	  }
// 	return &FirebaseClient{}

// }



// func c()  {

// 	proxyURL := "http://localhost:10100" // insert you proxy here 

// // The SDK makes 2 different types of calls:
// // 1. To the Google OAuth2 service to fetch the refresh and access tokens.
// // 2. To Firebase to send the pushes.
// // Each type uses its own HTTP Client and we need to insert our custom HTTP Client with proxy everywhere.
// cl := &http.Client{
// 	Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
// }
// ctxWithClient := context.WithValue(ctx, oauth2.HTTPClient, cl)

// // This is how we insert our custom HTTP Client in the Google OAuth2 service:
// // by context with specific value.
// creds, err := google.CredentialsFromJSON(ctxWithClient, []byte(c.Firebase.Credentials), firebaseScopes...)
// if err != nil {
//     log.Fatalf("google credentials from JSON: %s", err)
// }

// // And this is how we insert proxy for the Firebase calls. Initialize base transport with our proxy.
// tr := &oauth2.Transport{
// 	Source: creds.TokenSource,
// 	Base:   &http.Transport{Proxy: http.ProxyURL(proxyURL)},
// }

// hCl := &http.Client{
// 	Transport: tr,
// 	Timeout:   10 * time.Second,
// }

// opts := []option.ClientOption{option.WithHTTPClient(hCl)}

// app, err := firebase.NewApp(ctx, nil, opts...)
// if err != nil {
//     log.Fatalf("new firebase app: %s", err)
// }

// fcmClient, err := app.Messaging(context.TODO())
// if err != nil {
//     log.Fatalf("messaging: %s", err)
// }
	
// }