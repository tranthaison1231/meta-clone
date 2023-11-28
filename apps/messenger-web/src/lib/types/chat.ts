import type { Message } from './message';
import type { BasePaginationRequest } from './response';
import type { User } from './user';

export interface Chat {
	id: number;
	name?: string;
	members: User[];
	ownerId: string;
	lastMessage: Message;
	createdAt: string;
	updatedAt: string;
}

export interface CreateChatInputDto {
	memberIds: number[];
}

export interface GetChatsInputDto extends BasePaginationRequest {
	memberIds?: string[];
}
