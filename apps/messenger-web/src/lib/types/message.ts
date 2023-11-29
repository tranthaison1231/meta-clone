import type { BasePaginationRequest } from './response';

export interface Message {
	id: string;
	content: string;
	chatId: string;
	ownerId: string;
}

export type SendMessageInputDto = Omit<Message, 'id' | 'ownerId'>;

export interface GetMessagesInputDto extends BasePaginationRequest {
	targetMessageId?: string | undefined;
	isUp: boolean;
	chatId: string;
}
