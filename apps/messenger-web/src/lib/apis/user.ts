import { BASE_URL } from "$lib/services/request";
import { getToken } from "$lib/services/storage";
import type { AcceptFriendInputDto, AddFriendInputDto, GetUsersResponse } from "$lib/types";
import type { BaseResponseType } from "$lib/types/response";

export const usersApi = {
  getAll: async () => {
		const rest = await fetch(`${BASE_URL}/users`, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});
		return await rest.json() as BaseResponseType<GetUsersResponse> | undefined;
	},
  addFriend: async ({ friend_id,user_id }: AddFriendInputDto) => {
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
		return await rest.json() 
  },
	acceptFriend:async ({friend_id,user_id}:AcceptFriendInputDto) => {
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
		return await rest.json() 
	}
}