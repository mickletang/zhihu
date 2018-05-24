package models

import (
	"html/template"
)

type BasicUserinfo struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Gender    int    `json:"gender"`
	Headline  string `json:"headline"`
	AvatarURL string `json:"avatar_url"`
	URLToken  string `json:"url_token"`

	Posts         []*Post
	AnswerCount   uint `json:"answer_count"`
	FollowerCount uint `json:"follower_count"`

	Followed  bool `json:"is_followed"` //followed by user who sent request
	Following bool `json:"is_following"`
}

type Member struct {
	BasicUserinfo
	MarkedCount    uint `json:"marked_count"`
	FollowingCount uint `json:'following_count"`
	//	Educations []*Education
	PrivacyProtected bool `json:"is_privacy_protected"`
	VoteupCount      uint `json:"voteup_count"`
	//Blocked bool `json:"is_blocked"`
	Description         string `json:"description"`
	FollowingTopicCount uint   `json:"following_topic_count"`
	ThankedCount        uint   `json:"thanked_count"`
}

type AuthUserinfo struct {
	ID       uint
	Email    string //username
	Name     string
	Password string
}

type Follower struct {
	BasicUserinfo
}

type Voter struct {
	BasicUserinfo
}

type Author struct {
	Member
}

type Question struct {
	ID           string         `json:"id"`
	User         *BasicUserinfo `json:"user"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	DateCreated  string         `json:"date_created"`
	DateModified string         `json:"date_modified"`

	VisitCount    uint `json:"visit_count"`
	AnswerCount   uint `json:"answer_count"`
	CommentCount  uint `json:"comment_count"`
	FollowerCount uint `json:"follower_count"`

	Topics  []*Topic
	Answers []*Answer
	//	Comments  []*QuestionComment
	//	Followers []*User

	Followed             bool `json:"is_followed"` //followed by user who sent request
	Answered             bool `json:"is_answered"`
	VisitorAnswerID      uint `json:"visitor_answer_id"`
	VisitorAnswerDeleted bool `json:"visitor_answer_deleted"`
}

type Topic struct {
	ID   uint
	Name string
}

type Answer struct {
	ID string `json:"id"`
	*Question
	Author       *Author `json:"author"`
	Content      template.HTML
	DateCreated  string
	DateModified string

	MarkedCount   uint
	UpvoteCount   uint
	DownvoteCount uint
	CommentCount  uint

	Comments []*AnswerComment
	//	Upvotes  []*User

	Upvoted   bool
	Downvoted bool
	IsAuthor  bool
	Deleted   bool
}

type Post struct {
}

type Comment struct {
	ID            uint           `json:"id"`
	Author        *BasicUserinfo `json:"author"`
	DateCreated   string         `json:"date_created"`
	UpvoteCount   uint           `json:"upvote_count"`
	DownvoteCount uint           `json:"downvote_count"`
	Content       string         `json:"content"`
	LikeCount     uint           `json:"like_count"`
	Liked         bool           `json:"is_liked"`
}

type AnswerComment struct {
	*Answer
	Comment
}

type QuestionComment struct {
	*Question
	Comment
}

func NewQuestion() *Question {
	question := new(Question)
	question.User = new(BasicUserinfo)
	return question
}

func NewAnswer() *Answer {
	answer := new(Answer)
	answer.Question = NewQuestion()
	answer.Author = new(Author)
	return answer
}