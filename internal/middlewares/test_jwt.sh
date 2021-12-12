http -v --json POST localhost:8000/login username=admin password=admin


http -v -f GET localhost:8000/auth/refresh_token "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzYyMjExMTIsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTYzNjIxNzUxMn0.KPNkYlQKjwayW_62kusVt3Wvt33izumkLwGWBErB1iM"  "Content-Type: application/json"


http -f GET localhost:8000/auth/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzYyMjEzNTAsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTYzNjIxNzc1MH0.pVpM-zc2a5VswGLLnKC6Dibah2zjysReOaUqslA7GE8"  "Content-Type: application/json"
