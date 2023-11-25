import { request } from '$lib/services/request';

export const chatsApi = {
	getAll: async () => {
		const data = await request(`/chats`);
		return data;
	}
};
