package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizDifficulty string

const (
	QuizDifficultyEasy   QuizDifficulty = "Fácil"
	QuizDifficultyMedium QuizDifficulty = "Médio"
	QuizDifficultyHard   QuizDifficulty = "Difícil"
)

var ValidQuizDifficulties = map[QuizDifficulty]bool{
	QuizDifficultyEasy:   true,
	QuizDifficultyMedium: true,
	QuizDifficultyHard:   true,
}

type QuizQuestion struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"  json:"id"`
	Section     string             `bson:"section"        json:"section"`
	Topic       string             `bson:"topic"          json:"topic"`
	Question    string             `bson:"question"       json:"question"`
	Answer      string             `bson:"answer"         json:"answer"`
	Explanation string             `bson:"explanation"    json:"explanation"`
	Difficulty  QuizDifficulty     `bson:"difficulty"     json:"difficulty"`
	Tags        []string           `bson:"tags"           json:"tags"`
	CreatedAt   time.Time          `bson:"created_at"     json:"createdAt"`
}
