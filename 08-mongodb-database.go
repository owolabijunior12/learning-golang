package main

import (
	"fmt"
	"time"
)

// COURSE 8: MONGODB AND NOSQL DATABASES
// Topics covered:
// 1. MongoDB connection
// 2. BSON and document structure
// 3. CRUD operations
// 4. Filtering and querying
// 5. Aggregation pipeline
// 6. Indexes
// 7. Error handling
// 8. Best practices

// Note: Requires "go.mongodb.org/mongo-driver/mongo"

// ============ 1. DOCUMENT MODEL ============
type Product struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Price     float64   `bson:"price"`
	Category  string    `bson:"category"`
	InStock   bool      `bson:"inStock"`
	Tags      []string  `bson:"tags"`
	CreatedAt time.Time `bson:"createdAt"`
}

type Order struct {
	ID        string    `bson:"_id,omitempty"`
	UserID    string    `bson:"userId"`
	Products  []string  `bson:"products"`
	Total     float64   `bson:"total"`
	Status    string    `bson:"status"` // pending, shipped, delivered
	CreatedAt time.Time `bson:"createdAt"`
}

// ============ 2. CONNECTION PATTERN ============
// func connectMongo(uri string) (*mongo.Client, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
//	if err != nil {
//		return nil, err
//	}
//
//	err = client.Ping(ctx, nil)
//	return client, err
// }

// ============ 3. MONGODB OPERATIONS (Pseudo-code patterns) ============

// INSERT DOCUMENT
// func insertProduct(collection *mongo.Collection, product Product) (string, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	result, err := collection.InsertOne(ctx, product)
//	if err != nil {
//		return "", err
//	}
//
//	return result.InsertedID.(string), nil
// }

// INSERT MULTIPLE
// func insertProducts(collection *mongo.Collection, products []Product) ([]string, error) {
//	var docs []interface{}
//	for _, p := range products {
//		docs = append(docs, p)
//	}
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	result, err := collection.InsertMany(ctx, docs)
//	return result.InsertedIDs, err
// }

// FIND ONE DOCUMENT
// func findProductByName(collection *mongo.Collection, name string) (*Product, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	var product Product
//	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&product)
//	return &product, err
// }

// FIND MULTIPLE DOCUMENTS
// func findByCategory(collection *mongo.Collection, category string) ([]Product, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	cursor, err := collection.Find(ctx, bson.M{"category": category})
//	if err != nil {
//		return nil, err
//	}
//	defer cursor.Close(ctx)
//
//	var products []Product
//	if err = cursor.All(ctx, &products); err != nil {
//		return nil, err
//	}
//
//	return products, nil
// }

// UPDATE DOCUMENT
// func updateProduct(collection *mongo.Collection, id string, update bson.M) error {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	result, err := collection.UpdateOne(
//		ctx,
//		bson.M{"_id": id},
//		bson.M{"$set": update},
//	)
//
//	if result.MatchedCount == 0 {
//		return fmt.Errorf("document not found")
//	}
//
//	return err
// }

// DELETE DOCUMENT
// func deleteProduct(collection *mongo.Collection, id string) error {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
//	if result.DeletedCount == 0 {
//		return fmt.Errorf("document not found")
//	}
//
//	return err
// }

// ============ 4. ADVANCED QUERIES ============

// FILTERING
// findExpensive := collection.Find(ctx, bson.M{
//	"price": bson.M{"$gt": 100}, // greater than
// })

// COMPLEX FILTERS
// var opts []*options.FindOptions
// opts = append(opts, options.Find().SetLimit(10))
// opts = append(opts, options.Find().SetSort(bson.M{"price": -1}))
//
// cursor, err := collection.Find(ctx, bson.M{
//	"$or": []bson.M{
//		{"category": "electronics"},
//		{"category": "books"},
//	},
// }, opts...)

// ============ 5. AGGREGATION PIPELINE ============

// func aggregateProducts(collection *mongo.Collection) ([]bson.M, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	pipeline := mongo.Pipeline{
//		// Stage 1: Match documents
//		bson.D{{Key: "$match", Value: bson.D{{Key: "price", Value: bson.D{{Key: "$gte", Value: 50}}}}}},
//
//		// Stage 2: Group by category
//		bson.D{{Key: "$group", Value: bson.D{
//			{Key: "_id", Value: "$category"},
//			{Key: "avgPrice", Value: bson.D{{Key: "$avg", Value: "$price"}}},
//			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
//		}}},
//
//		// Stage 3: Sort by count descending
//		bson.D{{Key: "$sort", Value: bson.D{{Key: "count", Value: -1}}}},
//	}
//
//	cursor, err := collection.Aggregate(ctx, pipeline)
//	if err != nil {
//		return nil, err
//	}
//	defer cursor.Close(ctx)
//
//	var results []bson.M
//	if err = cursor.All(ctx, &results); err != nil {
//		return nil, err
//	}
//
//	return results, nil
// }

// ============ 6. INDEXES ============

// func createIndexes(collection *mongo.Collection) error {
//	indexModel := mongo.IndexModel{
//		Keys: bson.D{{Key: "name", Value: 1}},
//	}
//
//	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
//	return err
// }

// ============ COURSE EIGHT MAIN FUNCTION ============
func courseEight() {
	fmt.Println("=== MONGODB AND NOSQL DATABASES ===\n")

	fmt.Println("MONGODB SETUP:")
	fmt.Println("---\n")

	fmt.Println("Docker MongoDB:")
	fmt.Println(`docker run --name mongodb -d -p 27017:27017 mongo:latest`)
	fmt.Println()

	fmt.Println("Connection String:")
	fmt.Println(`mongodb://localhost:27017`)
	fmt.Println()

	fmt.Println("BASIC CRUD CODE PATTERN:")
	fmt.Println("---\n")

	fmt.Println(`
// Connect
client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
defer client.Disconnect(ctx)

// Get collection
collection := client.Database("mydb").Collection("products")

// INSERT
result, err := collection.InsertOne(ctx, Product{
	Name:  "Laptop",
	Price: 999.99,
	Tags:  []string{"electronics", "computers"},
})

// FIND ONE
var product Product
err := collection.FindOne(ctx, bson.M{"name": "Laptop"}).Decode(&product)

// FIND MULTIPLE
cursor, err := collection.Find(ctx, bson.M{"price": bson.M{"$gt": 100}})
var products []Product
cursor.All(ctx, &products)

// UPDATE
result, err := collection.UpdateOne(
	ctx,
	bson.M{"_id": id},
	bson.M{"$set": bson.M{"price": 1099.99}},
)

// DELETE
result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
`)
	fmt.Println()

	fmt.Println("MONGODB QUERY OPERATORS:")
	fmt.Println("---")
	fmt.Println("$eq    - Equal")
	fmt.Println("$gt    - Greater than")
	fmt.Println("$gte   - Greater than or equal")
	fmt.Println("$lt    - Less than")
	fmt.Println("$lte   - Less than or equal")
	fmt.Println("$ne    - Not equal")
	fmt.Println("$in    - In array")
	fmt.Println("$nin   - Not in array")
	fmt.Println("$and   - Logical AND")
	fmt.Println("$or    - Logical OR")
	fmt.Println("$not   - Logical NOT")
	fmt.Println("$exists - Field exists")
	fmt.Println("$regex - Regular expression")
	fmt.Println()

	fmt.Println("UPDATE OPERATORS:")
	fmt.Println("---")
	fmt.Println("$set       - Set field value")
	fmt.Println("$unset     - Remove field")
	fmt.Println("$inc       - Increment value")
	fmt.Println("$push      - Add to array")
	fmt.Println("$pull      - Remove from array")
	fmt.Println("$addToSet  - Add to set (if not exists)")
	fmt.Println("$rename    - Rename field")
	fmt.Println("$currentDate - Set to current date")
	fmt.Println()

	fmt.Println("AGGREGATION PIPELINE STAGES:")
	fmt.Println("---")
	fmt.Println("$match     - Filter documents (like WHERE)")
	fmt.Println("$group     - Group and aggregate")
	fmt.Println("$sort      - Sort documents")
	fmt.Println("$limit     - Limit result count")
	fmt.Println("$skip      - Skip documents")
	fmt.Println("$project   - Select fields (like SELECT)")
	fmt.Println("$lookup    - JOIN with other collection")
	fmt.Println("$unwind    - Expand array fields")
	fmt.Println("$count     - Count documents")
	fmt.Println()

	fmt.Println("AGGREGATION EXAMPLE:")
	fmt.Println("---")
	fmt.Println(`
pipeline := mongo.Pipeline{
	// Find products over $100
	bson.D{{Key: "$match", Value: bson.D{
		{Key: "price", Value: bson.D{{Key: "$gt", Value: 100}}},
	}}},
	
	// Group by category and calculate stats
	bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: "$category"},
		{Key: "avgPrice", Value: bson.D{{Key: "$avg", Value: "$price"}}},
		{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		{Key: "maxPrice", Value: bson.D{{Key: "$max", Value: "$price"}}},
	}}},
	
	// Sort by count descending
	bson.D{{Key: "$sort", Value: bson.D{{Key: "count", Value: -1}}}},
}

cursor, _ := collection.Aggregate(ctx, pipeline)
`)
	fmt.Println()

	fmt.Println("INDEXING:")
	fmt.Println("---")
	fmt.Println(`
// Create index on name field
indexModel := mongo.IndexModel{
	Keys: bson.D{{Key: "name", Value: 1}},
}
collection.Indexes().CreateOne(ctx, indexModel)

// Compound index
indexModel := mongo.IndexModel{
	Keys: bson.D{
		{Key: "category", Value: 1},
		{Key: "price", Value: -1},
	},
}

// Unique index
opts := options.Index().SetUnique(true)
indexModel := mongo.IndexModel{
	Keys:    bson.D{{Key: "email", Value: 1}},
	Options: opts,
}
`)
	fmt.Println()

	fmt.Println("TRANSACTIONS:")
	fmt.Println("---")
	fmt.Println(`
session, err := client.StartSession()
defer session.EndSession(ctx)

err = session.WithTransaction(ctx, func(sc context.Context) error {
	// Operations 1
	collection.InsertOne(sc, order)
	
	// Operation 2
	collection.UpdateOne(sc, filter, update)
	
	// All succeed or all rollback
	return nil
})
`)
	fmt.Println()

	fmt.Println("BEST PRACTICES:")
	fmt.Println("---")
	fmt.Println("✓ Always use context with timeout")
	fmt.Println("✓ Close cursors after use")
	fmt.Println("✓ Use indexes on frequently queried fields")
	fmt.Println("✓ Validate data before inserting")
	fmt.Println("✓ Handle not found errors explicitly")
	fmt.Println("✓ Use aggregation for complex queries")
	fmt.Println("✓ Structure documents efficiently")
	fmt.Println("✓ Monitor query performance")
	fmt.Println("✓ Use transactions for related operations")
	fmt.Println("✓ Batch operations when possible")
	fmt.Println()

	fmt.Println("COMMON LIBRARIES:")
	fmt.Println("---")
	fmt.Println("go.mongodb.org/mongo-driver    - Official MongoDB driver")
	fmt.Println("github.com/qiniu/qmgo          - Wrapper around mongo driver")
	fmt.Println("entgo.io/ent                   - Entity framework (supports MongoDB)")
	fmt.Println()

	fmt.Println("=== END OF MONGODB ===")
}

// KEY TAKEAWAYS:
// 1. MongoDB stores JSON-like documents (BSON)
// 2. No schema required - flexible document structure
// 3. _id field is required and must be unique
// 4. Always use context with timeout
// 5. FindOne returns single document, Find returns cursor
// 6. Close cursor after reading to free resources
// 7. Decode converts BSON to Go struct
// 8. UpdateOne uses $set operator to set fields
// 9. MatchedCount tells if document was found
// 10. ModifiedCount tells how many were modified
// 11. DeletedCount tells how many were deleted
// 12. Aggregation pipeline is powerful for complex queries
// 13. Indexes significantly improve query performance
// 14. Transactions ensure data consistency
// 15. Use batch operations for multiple inserts
// 16. Array fields can be queried and updated
// 17. Regular expressions for flexible text search
// 18. TTL indexes can auto-delete old documents
// 19. Validation rules can be set at collection level
// 20. MongoDB is great for flexible, document-oriented data
