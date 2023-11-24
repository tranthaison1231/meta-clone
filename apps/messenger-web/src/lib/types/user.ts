export interface User {
	email: string;
	name: string;
	id: number;
	avatar: string;
	created_at: string;
	updated_at: string;
	gender: string;
	friends: User[] | null;
	friend_requests: FriendRequest[] | null
}

export interface FriendRequest {
	id: number;
	user_id: number;
	friend_id: number
}

export enum FriendStatus {
	Pending = 'Pending',
	RequireAccept = 'RequireAccept',
	Friend = 'Friend',
	None = 'None'
}

export interface GetUsersResponse {
  users: User & { friend_status: FriendStatus }[]
}

export interface AddFriendInputDto {
	user_id: number;
	friend_id: number;
}

export type AcceptFriendInputDto = AddFriendInputDto
