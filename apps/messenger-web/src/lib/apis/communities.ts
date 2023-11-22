import { BASE_URL } from '$lib/services/request';

export const communitiesApi = {
	getAll: async () => {
		const rest = await fetch(`${BASE_URL}/communities`, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});
		return rest.json();
	}
};
