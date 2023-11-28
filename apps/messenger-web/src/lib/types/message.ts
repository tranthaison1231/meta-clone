import type { BasePaginationRequest } from "./response";

export interface Message {
  id: number;
  content: string
  chatId: number;
  ownerId: number
}

export interface SendMessageInputDto extends Omit<Message, "id" | "ownerId"> {
  chatId: number
}

export interface GetMessagesInputDto extends BasePaginationRequest {
  targetMessageId?: number | undefined
  isUp: boolean
  chatId: number
}