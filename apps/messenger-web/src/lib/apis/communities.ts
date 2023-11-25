import { request } from '$lib/services/request';

export const communitiesApi = {
	getAll: async () => {
		const data = await request(`/communities`);
		return data;
	}
};
