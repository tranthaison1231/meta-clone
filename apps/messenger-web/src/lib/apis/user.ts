import { DEFAULT_PAGE, DEFAULT_PAGE_SIZE } from '$lib/constants/pagination';
import { BASE_URL, request } from '$lib/services/request';
import { getToken } from '$lib/services/storage';
import type {
	AcceptFriendInputDto,
	AddFriendInputDto,
	GetUserFriendsInputDto,
	GetUsersResponse,
	User
} from '$lib/types';
import type {
	BasePaginationRequest,
	BasePaginationResponse,
	BaseResponseType
} from '$lib/types/response';

export const usersApi = {
	getAll: async ({
		limit = DEFAULT_PAGE_SIZE,
		page = DEFAULT_PAGE,
		orderBy
	}: BasePaginationRequest) => {
		const searchParams = new URLSearchParams({
			limit: String(limit),
			page: String(page),
			...(orderBy
				? {
						order_by: String(orderBy)
				  }
				: {})
		});
		const { data } = await request.get<
			unknown,
			BaseResponseType<BasePaginationResponse<GetUsersResponse>>
		>(`${BASE_URL}/users?${searchParams}`);

		return data;
	},
	getFriends: async ({
		limit = DEFAULT_PAGE_SIZE,
		page = DEFAULT_PAGE,
		orderBy,
		userId
	}: GetUserFriendsInputDto) => {
		const searchParams = new URLSearchParams({
			limit: String(limit),
			page: String(page),
			...(orderBy
				? {
						order_by: String(orderBy)
				  }
				: {})
		});
		const { data } = await request.get<unknown, BaseResponseType<BasePaginationResponse<User>>>(
			`${BASE_URL}/users/${userId}/friends?${searchParams}`
		);
		return data;
	},
	addFriend: async ({ friend_id, user_id }: AddFriendInputDto) => {
		const rest = await fetch(`${BASE_URL}/users/add-friend`, {
			method: 'POST',
			body: JSON.stringify({
				friend_id,
				user_id
			}),
			headers: {
				Authorization: `Bearer ${getToken()}`
			}
		});
		return await rest.json();
	},
	acceptFriend: async ({ friend_id, user_id }: AcceptFriendInputDto) => {
		const rest = await fetch(`${BASE_URL}/users/accept-friend`, {
			method: 'POST',
			body: JSON.stringify({
				friend_id,
				user_id
			}),
			headers: {
				Authorization: `Bearer ${getToken()}`
			}
		});
		return await rest.json();
	}
};
