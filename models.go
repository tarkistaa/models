package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subject struct {
	Id         primitive.ObjectID   `bson:"_id",json:"id"`
	Name       string               `bson:"name",json:"name"`
	AuthorId   primitive.ObjectID   `bson:"authorId",json:"authorId"`
	TeacherIds []primitive.ObjectID `bson:"teacherIds",json:"teacherIds"`
	Themes     []Theme              `bson:"themes",json:"themes"`
	Tests      []primitive.ObjectID `bson:"testIds",json:"testIds"`
}

type Theme struct {
	Id   primitive.ObjectID `bson:"_id",json:"id"`
	Name string             `bson:"name",json:"name"`
}

type Test struct {
	Id          primitive.ObjectID   `bson:"_id",json:"id"`
	Name        string               `bson:"name",json:"name"`
	AuthorId    primitive.ObjectID   `bson:"authorId",json:"authorId"`
	SubjectId   primitive.ObjectID   `bson:"subjectId",json:"subjectId"`
	ThemeId     primitive.ObjectID   `bson:"themeId",json:"themeId"`
	QuestionIds []primitive.ObjectID `bson:"questionIds",json:"questionIds"`
	TimeLimited bool                 `bson:"timeLimited",json:"timeLimited"`
	TimeLimit   int32                `bson:"timeLimit",json:"timeLimit"`
}

type Question struct {
	Id       primitive.ObjectID `bson:"_id",json:"id"`
	TestId   primitive.ObjectID `bson:"testId",json:"testId"`
	Question string             `bson:"question",json:"question"`
	Options  []Option           `bson:"options",json:"options"`
}

type Option struct {
	Id      primitive.ObjectID `bson:"_id",json:"id"`
	Option  string             `bson:"option",json:"option"`
	IsRight bool               `bson:"isRight",json:"isRight"`
}

type User struct {
	Id                primitive.ObjectID   `bson:"_id",json:"id"`
	FirstName         string               `bson:"firstName",json:"firstName"`
	SecondName        string               `bson:"secondName",json:"secondName"`
	Patronymic        string               `bson:"patronymic",json:"patronymic"`
	Login             string               `bson:"login",json:"login"`
	Password          string               `bson:"password",json:"password"`
	Type              string               `bson:"type",json:"type"`
	SubjectIds        []primitive.ObjectID `bson:"subjectIds",json:"subjectIds"`
	CreatedSubjectIds []primitive.ObjectID `bson:"createdSubjectIds",json:"createdSubjectIds"`
}
