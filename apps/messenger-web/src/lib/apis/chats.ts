import { DEFAULT_PAGE, DEFAULT_PAGE_SIZE } from '$lib/constants/pagination';
import { request } from '$lib/services/request';
import type { Chat, CreateChatInputDto, GetChatsInputDto } from '$lib/types/chat';
import type { BasePaginationResponse } from '$lib/types/response';

export const chatsApi = {
	getAll: async ({
		memberIds = [],
		limit = DEFAULT_PAGE_SIZE,
		orderBy,
		page = DEFAULT_PAGE,
		isSingleChat
	}: GetChatsInputDto) => {
		const searchParams = new URLSearchParams({
			limit: String(limit),
			page: String(page),
			memberIds: memberIds.join(',')
		});

		if (isSingleChat) {
			searchParams.append('isSingleChat', String(isSingleChat));
		}

		if (orderBy) {
			searchParams.append('orderBy', orderBy);
		}

		const { data } = await request(`/chats?${searchParams}`);

		return data as BasePaginationResponse<Chat>;
	},
	createChat: async ({ memberIds }: CreateChatInputDto) => {
		const { data } = await request('/chats', {
			method: 'POST',
			body: JSON.stringify({
				memberIds
			})
		});
		return data as { chat: Chat };
	}
};
