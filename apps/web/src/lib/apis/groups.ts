import { BASE_URL } from '$lib/services/request';

export const groupsApi = {
	getAll: async () => {
		const rest = await fetch(`${BASE_URL}/me/groups`);
		return rest.json();
	}
};
