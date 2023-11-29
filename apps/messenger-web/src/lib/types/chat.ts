import type { Message } from './message';
import type { BasePaginationRequest } from './response';
import type { User } from './user';

export interface Chat {
	id: string;
	name?: string;
	members: User[];
	ownerId: string;
	lastMessage: Message;
	createdAt: string;
	updatedAt: string;
}

export interface CreateChatInputDto {
	memberIds: string[];
}

export interface GetChatsInputDto extends BasePaginationRequest {
	memberIds?: string[];
	isSingleChat?: boolean;
}
