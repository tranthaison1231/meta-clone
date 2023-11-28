import { DEFAULT_PAGE, DEFAULT_PAGE_SIZE } from '$lib/constants/pagination';
import { request } from '$lib/services/request';
import type { GetMessagesInputDto, Message, SendMessageInputDto } from '$lib/types/message';
import type { BasePaginationResponse } from '$lib/types/response';

export const messagesApi = {
	sendMessage: async ({ chatId, content }: SendMessageInputDto) => {
		const { data } = await request(`/chats/${chatId}/messages`, {
			method: 'POST',
			body: JSON.stringify({
				content
			})
		});
		return data as { message: Message };
	},
	getMessages: async ({
		limit = DEFAULT_PAGE_SIZE,
		page = DEFAULT_PAGE,
		orderBy,
		targetMessageId,
		isUp,
		chatId
	}: GetMessagesInputDto) => {
		const searchParams = new URLSearchParams({
			limit: String(limit),
			page: String(page),
			isUp: String(isUp),
		});

		if (targetMessageId) {
			searchParams.append('targetMessageId', String(targetMessageId));
		}

		if (orderBy) {
			searchParams.append('orderBy', orderBy);
		}

		const { data } = await request(`/chats/${chatId}/messages?${searchParams}`);
		return data as BasePaginationResponse<Message>;
	}
};
