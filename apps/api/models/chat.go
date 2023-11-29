package models

type Chat struct {
	Base
	Name          string   `gorm:"type:varchar(255)" json:"name"`
	Members       []*User  `gorm:"many2many:chat_users;" json:"members"`
	OwnerID       string   `json:"ownerId"`
	Owner         User     `gorm:"foreignKey:OwnerID;references:ID" json:"owner"`
	LastMessageID string   `json:"lastMessageId"`
	LastMessage   *Message `gorm:"foreignKey:LastMessageID;references:ID" json:"lastMessage"`
}

type ChatUser struct {
	ChatID string `json:"chatId"`
	UserID string `json:"userId"`
}

type CreateChatRequest struct {
	MemberIDs []string `json:"memberIds"`
}

type GetChatByMemberIdsRequest struct {
	CreateChatRequest
}

type AddMemberToChatRequest struct {
	UserID string `json:"userId" validate:"required"`
}

type GetChatsRequest struct {
	PaginateRequest BasePaginationRequest `json:"paginateRequest"`
	MemberIds       []string              `json:"memberIds"`
	IsSingleChat    bool                  `json:"isSingleChat"`
}

type GetChatMessagesRequest struct {
	PaginateRequest BasePaginationRequest `json:"paginateRequest"`
	ChatID          string                `json:"chatId"`
	IsUp            bool                  `json:"isUp"`
	TargetMessageID string                `json:"targetMessageId"`
}
