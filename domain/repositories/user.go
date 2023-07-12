package repositories

import (
	. "bn-login-repositories/domain/datasources"
	"bn-login-repositories/domain/entities"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

type IUserRepository interface {
	GetVoiceStudio(user_id string) *entities.SpeakerModel
	// getAllBlacklist() ([]bson.M, error)
	// getAllUsers() ([]bson.M, error)
	CheckBlacklist() *entities.LineLoginData
	// getAllUserData(userID string) (bson.M, error)
	// findEmail(email string) (bson.M, error)
	// getEmail(email string) string
	// getExistingJWT(userID, jwt string) bson.M
	// getUserID(userID string) string
	// insertNewUser(dataJSON bson.M) error
	// cleaningUserData(userID string) error
	// updateSignInProvider(userID, signInProvider string)
	// updateJWT(userID string, packageData map[string]interface{}) error
	// updateLogin(userID, token string) error
	// updateProfile(userID, picture, email string) error
	// updateEmail(userID, email string) error
	// updatePlatformData(userData bson.M, packageData map[string]interface{})
	// updateProfilePicture(userData bson.M, lastLoginType string) error
	// changeNaNToList(userID string) error
	// removePlatform(userID string) error
	// updateAgreement(data bson.M) error
	// updateAgreementFalse(data bson.M) error
	// updatePersonalFalse(data bson.M) error
	// updateAccountStatus(data bson.M, status string) error
	// updateUUID(data bson.M, uuid string) error
	// updateSubscription(data bson.M) error
	// insertPlayQuota(data bson.M) error
	// createNewName(userID, name string) error
	// checkName(name string) string
	// FindPrivateSpeaker() ([]entities.SpeakerModel, error)
	// FindPublicSpeaker() ([]entities.SpeakerModel, error)
}

func NewUserRepository(db *MongoDB) IUserRepository {
	return &userRepository{
		Context:    db.Context,
		Collection: db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("user"),
	}
}

func (repo userRepository) GetVoiceStudio(user_id string) *entities.SpeakerModel {
	var result entities.SpeakerModel
	filter := bson.M{"user_id": user_id}
	err := repo.Collection.FindOne(repo.Context, filter).Decode(&result)
	if err != nil {
		fmt.Println("No matching document found")
		return nil
	}
	return &result
}

// func (repo userRepository) getAllBlacklist() ([]bson.M, error) {
// 	filter := bson.M{}
// 	var blacklist []bson.M
// 	cursor, err := repo.Collection.Find(repo.Context, filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(repo.Context)
// 	for cursor.Next(repo.Context) {
// 		var result bson.M
// 		err := cursor.Decode(&result)
// 		if err != nil {
// 			return nil, err
// 		}
// 		blacklist = append(blacklist, result)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}
// 	return blacklist, nil
// }

// func (u *userRepository) getAllUsers() ([]bson.M, error) {
// 	filter := bson.M{}
// 	var users []bson.M
// 	cur, err := u.Collection.Find(u.Context, filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cur.Close(u.Context)
// 	for cur.Next(u.Context) {
// 		var result bson.M
// 		err := cur.Decode(&result)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, result)
// 	}
// 	if err := cur.Err(); err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

func (repo userRepository) CheckBlacklist(userID string) error {
	options := options.Find()
	filter := bson.M{"user_id": userID}
	cursor, err := repo.Collection.Find(repo.Context, filter, options)

	if err != nil {
		return "Failed"
	}
	fmt.Println(cursor)
	return "OK"
}

// func (u *userRepository) getAllUserData(userID string) (bson.M, error) {
// 	filter := bson.M{"user_id": userID}
// 	var result bson.M
// 	err := usersCollection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (u *UserRepository) findEmail(email string) (bson.M, error) {
// 	filter := bson.M{"email": email}
// 	var result bson.M
// 	err := usersCollection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (u *UserRepository) getEmail(email string) string {
// 	filter := bson.M{"email": email}
// 	var result bson.M
// 	err := usersCollection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		return ""
// 	}
// 	return result["email"].(string)
// }

// func (u *userRepository) getUserID(userID string) string {
// 	filter := bson.M{"user_id": userID}
// 	var result bson.M
// 	err := usersCollection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		return ""
// 	}
// 	return result["user_id"].(string)
// }

// func (u *userRepository) getExistingJWT(userID, jwt string) bson.M {
// 	filter := bson.M{"user_id": userID, "jwt": jwt}
// 	var result bson.M
// 	err := usersCollection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		return nil
// 	}
// 	return result
// }

// func (u *userRepository) insertNewUser(dataJSON bson.M) error {
// 	_, err := usersCollection.InsertOne(ctx, dataJSON)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) cleaningUserData(userID string) error {
// 	filter := bson.M{"user_id": userID}
// 	update := bson.M{"$unset": bson.M{
// 		"last_login_type": 1,
// 		"platform":        1,
// 	}}
// 	_, err := usersCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) updateSignInProvider(userID, signInProvider string) error {
// 	filter := bson.M{"user_id": userID}
// 	update := bson.M{"$set": bson.M{
// 		"sign_in_provider": signInProvider,
// 	}}
// 	_, err := usersCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) updateJWT(userID string, packageData map[string]interface{}) error {
// 	filter := bson.M{"user_id": userID}
// 	update := bson.M{"$set": bson.M{
// 		"end_datetime":     time.Now(),
// 		"jwt":              packageData["token"],
// 		"image":            packageData["image"],
// 		"last_login_type":  packageData["last_login_type"],
// 	}}
// 	_, err := usersCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) updateLogin(userID, token string) error {
// 	filter := bson.M{"user_id": userID}
// 	update := bson.M{"$set": bson.M{
// 		"end_datetime": time.Now(),
// 		"jwt":          token,
// 	}}
// 	_, err := usersCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) updateProfile(userID, picture, email string) error {
// 	filter := bson.M{"user_id": userID}
// 	update := bson.M{"$set": bson.M{
// 		"image": picture,
// 		"email": email,
// 	}}
// 	_, err := usersCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) updateEmail(userID, email string) error {
// 	filter := bson.M{"user_id": userID}
// 	update := bson.M{"$set": bson.M{
// 		"email": email,
// 	}}
// 	_, err := usersCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) updatePlatformData(userData bson.M, packageData map[string]interface{}) error {
// 	filter := bson.M{"user_id": userData["user_id"]}
// 	update := bson.M{"$push": bson.M{
// 		"platform": packageData,
// 	}}
// 	_, err := usersCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) updateProfilePicture(userData bson.M, lastLoginType string) error {
// 	pic := ""
// 	if _, ok := userData["picture"]; ok {
// 		pic = userData["picture"].(string)
// 	} else if _, ok := userData["image"]; ok {
// 		pic = userData["image"].(string)
// 	} else {
// 		pic = ENV.DEFAULT_PICTURE
// 	}

// 	filter := bson.M{"user_id": userData["user_id"]}
// 	update := bson.M{"$set": bson.M{
// 		"username":         userData["username"],
// 		"image":            pic,
// 		"last_login_type":  lastLoginType,
// 	}}
// 	_, err := usersCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) changeNaNToList(userID string) error {
// 	filter := bson.M{"user_id": userID}
// 	update := bson.M{"$set": bson.M{
// 		"platform": []bson.M{{"line": true}},
// 	}}
// 	_, err := usersCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) removePlatform(userID string) error {
// 	filter := bson.M{"user_id": userID}
// 	update := bson.M{"$unset": bson.M{
// 		"platform": 1,
// 	}}
// 	_, err := usersCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) updateAgreement(data bson.M) error {
// 	if _, ok := data["agreement"]; !ok {
// 		_, err := usersCollection.UpdateOne(ctx, bson.M{"user_id": data["user_id"]}, bson.M{"$set": bson.M{"agreement": false}})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	_, err := usersCollection.UpdateOne(ctx, bson.M{"user_id": data["user_id"]}, bson.M{"$set": bson.M{"agreement": true}})
// 	if err != nil {
// 		return err
// 	}

// 	temp := bson.M{
// 		"user_id":   data["user_id"],
// 		"date_time": time.Now(),
// 	}
// 	_, err = logsAgreementCollection.InsertOne(ctx, temp)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) updateAgreementFalse(data bson.M) error {
// 	if _, ok := data["agreement"]; !ok {
// 		_, err := usersCollection.UpdateOne(ctx, bson.M{"user_id": data["user_id"]}, bson.M{"$set": bson.M{"agreement": false}})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (u *userRepository) updatePersonalFalse(data bson.M) error {
// 	if _, ok := data["personal_form"]; !ok {
// 		_, err := usersCollection.UpdateOne(ctx, bson.M{"user_id": data["user_id"]}, bson.M{"$set": bson.M{"personal_form": false}})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (u *userRepository) updateAccountStatus(data bson.M, status string) error {
// 	if _, ok := data["account_status"]; !ok {
// 		_, err := usersCollection.UpdateOne(ctx, bson.M{"user_id": data["user_id"]}, bson.M{"$set": bson.M{"account_status": status}})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (u *userRepository) updateUUID(data bson.M, uuid string) error {
// 	if _, ok := data["uid"]; !ok {
// 		_, err := usersCollection.UpdateOne(ctx, bson.M{"user_id": data["user_id"]}, bson.M{"$set": bson.M{"uid": uuid}})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (u *userRepository) updateSubscription(data bson.M) error {
// 	if _, ok := data["subscription"]; !ok {
// 		_, err := usersCollection.UpdateOne(ctx, bson.M{"user_id": data["user_id"]}, bson.M{"$set": bson.M{"subscription": "Free"}})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (u *userRepository) insertPlayQuota(data bson.M) error {
// 	if _, ok := data["play_quota"]; !ok {
// 		_, err := usersCollection.UpdateOne(ctx, bson.M{"user_id": data["user_id"]}, bson.M{"$set": bson.M{"play_quota": 2500}})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (u *userRepository) createNewName(userID, name string) error {
// 	_, err := usersCollection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$set": bson.M{"name": name}})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *userRepository) checkName(name string) string {
// 	filter := bson.M{"name": name}
// 	var result bson.M
// 	err := usersCollection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		return ""
// 	}
// 	return result["user_id"].(string)
// }
