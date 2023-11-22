import { BASE_URL } from '$lib/services/request';

export const chatsApi = {
	getAll: async () => {
		const rest = await fetch(`${BASE_URL}/chats`, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});
		return rest.json();
	}
};
