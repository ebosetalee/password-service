package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ebosetalee/password-service.git/types"
	"github.com/ebosetalee/password-service.git/utils"
	"github.com/golang-jwt/jwt/v5"
)

func WithJWTAuth(handlerFunc http.HandlerFunc, repo types.UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromRequest(r)

		token, err := verifyJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userID"].(string)

		user, err := repo.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		fmt.Println(user)
		// Add the user to the context
		// ctx := r.Context()
		// ctx = context.WithValue(ctx, UserKey, u.ID)
		// r = r.WithContext(ctx)

		// Call the function if the token is valid
		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("session expired, please log in again"))
}
