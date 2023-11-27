import { DEFAULT_PAGE, DEFAULT_PAGE_SIZE } from '$lib/constants/pagination';
import { request } from '$lib/services/request';
import type {
	AcceptFriendInputDto,
	AddFriendInputDto,
	GetFriendsResponse,
	GetUserFriendsInputDto,
	GetUsersResponse
} from '$lib/types';
import type { BasePaginationRequest, BasePaginationResponse } from '$lib/types/response';

export const usersApi = {
	getAll: async ({
		limit = DEFAULT_PAGE_SIZE,
		page = DEFAULT_PAGE,
		orderBy
	}: BasePaginationRequest) => {
		const searchParams = new URLSearchParams({
			limit: String(limit),
			page: String(page)
		});
		if (orderBy) {
			searchParams.append('order_by', orderBy);
		}
		const { data } = await request(`/users?${searchParams}`);
		return data as BasePaginationResponse<GetUsersResponse>;
	},
	getFriends: async ({
		limit = DEFAULT_PAGE_SIZE,
		page = DEFAULT_PAGE,
		orderBy,
		userId
	}: GetUserFriendsInputDto) => {
		const searchParams = new URLSearchParams({
			limit: String(limit),
			page: String(page)
		});
		if (orderBy) {
			searchParams.append('order_by', orderBy);
		}
		const { data } = await request(`/users/${userId}/friends?${searchParams}`);
		return data as BasePaginationResponse<GetFriendsResponse>;
	},
	addFriend: async ({ friendId, userId }: AddFriendInputDto) => {
		const data = await request(`/users/add-friend`, {
			method: 'POST',
			body: JSON.stringify({
				friendId,
				userId
			})
		});
		return data;
	},
	acceptFriend: async ({ friendId, userId, isRejecting }: AcceptFriendInputDto) => {
		const data = await request('/users/accept-friend', {
			method: 'POST',
			body: JSON.stringify({
				friendId,
				userId,
				isRejecting
			})
		});
		return data;
	}
};
