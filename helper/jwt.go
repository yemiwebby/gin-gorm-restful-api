package helper

// var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

// func GenerateJWT(user model.User) (string, error) {
// 	jwtSecret := []byte(os.Getenv("JWT_PRIVATE_KEY"))
// 	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id":  user.ID,
// 		"iat": time.Now().Unix(),
// 		// using "eat" doesn't actually expire the token
// 		"exp": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
// 	})
// 	return token.SignedString(jwtSecret)
// }
//
// func ValidateJWT(context *gin.Context) (*jwt.Token, error) {
// 	reqToken := getTokenFromRequest(context)
//
// 	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
// 		}
//
// 		jwtSecret := []byte(os.Getenv("JWT_PRIVATE_KEY"))
//
// 		return jwtSecret, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	_, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		return token, nil
// 	}
// 	return nil, errors.New("invalid token provided")
// }

// func CurrentUser(context *gin.Context) (model.User, error) {
// 	err := ValidateJWT(context)
// 	if err != nil {
// 		return model.User{}, err
// 	}
//
// 	token, _ := getToken(context)
// 	claims, _ := token.Claims.(jwt.MapClaims)
// 	userId := uint(claims["id"].(float64))
//
// 	user, err := model.FindUserById(userId)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return user, nil
// }

// func getToken(context *gin.Context) (*jwt.Token, error) {
// 	tokenString := getTokenFromRequest(context)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
//
// 		return privateKey, nil
// 	})
// 	return token, err
// }

// func getTokenFromRequest(context *gin.Context) string {
// 	bearerToken := context.Request.Header.Get("Authorization")
// 	splitToken := strings.Split(bearerToken, " ")
// 	if len(splitToken) == 2 {
// 		return splitToken[1]
// 	}
//
// 	// Get jwt from cookie if exist
// 	ct, err := context.Cookie("access_token")
// 	if err != nil {
// 		return ct
// 	}
// 	return ""
// }
