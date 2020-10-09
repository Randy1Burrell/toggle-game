package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/randy1burrell/toggle-game/data/models"
	"github.com/randy1burrell/toggle-game/helpers"
)

func Insert(deck *models.Deck) (interface{}, error) {
	client, coll, ctx, err := helpers.GetClient()
	if err != nil {
		return nil, err
	}

	defer client.Disconnect(*ctx)

	deckResult, err := coll.InsertOne(context.TODO(), deck)

	if err != nil {
		return nil, err
	}

	if oid, ok := deckResult.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return deckResult.InsertedID, err
}

// Find one object
func FindOne(id string) (models.DeckId, error) {
	client, coll, ctx, err := helpers.GetClient()
	var res models.DeckId

	if err != nil {
		return res, err
	}

	// Close database when done
	defer client.Disconnect(*ctx)

	// Make sure mongo's object id match with this
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return res, err
	}

	err = coll.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}}).Decode(&res)

	if err != nil {
		return res, err
	}

	// res.Cards = cards
	return res, err
}

func DrawCard(id string, num int) ([]models.Card, error) {
	var res []models.Card
	if deck, err := FindOne(id); err == nil {
		// Initialise Client
		client, coll, ctx, err := helpers.GetClient()
		if err != nil {
			return res, err
		}

		// Close out client connection when done
		defer client.Disconnect(*ctx)

		// mongo's object id and id being searched for have to be similar for a match
		objID, e := primitive.ObjectIDFromHex(id)

		if e != nil {
			return res, e
		}

		filter := bson.M{
			"_id": bson.M{
				"$eq": objID,
			},
		}

		if num > deck.Remaining {
			deck.Remaining = 0
			res = deck.Cards
			deck.Cards = nil
		} else {
			res = deck.Cards[:num]
			deck.Remaining -= num
			deck.Cards = deck.Cards[num:]
		}

		// Filter for the first few cards and mark them as used

		update := bson.M{
			"$set": bson.M{
				"cards":     deck.Cards,
				"remaining": deck.Remaining,
			},
		}
		// We don't care about the results after it is updated
		_, err = coll.UpdateOne(*ctx, filter, update)

		return res, err
	} else {
		return res, err
	}
}
