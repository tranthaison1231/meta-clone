import type { BasePaginationRequest } from './response';

export interface User {
	email: string;
	name: string;
	id: number;
	avatar: string;
	created_at: string;
	updated_at: string;
	firstName: string;
	lastName: string;
	gender: string;
	friends: User[] | null;
	friendRequests: FriendRequest[] | null;
}

export interface FriendRequest {
	id: number;
	userId: number;
	friendId: number;
}

export enum FriendStatus {
	Pending = 'Pending',
	RequireAccept = 'RequireAccept',
	Friend = 'Friend',
	None = 'None'
}

export interface GetUsersResponse {
	users: (User & { friendStatus: FriendStatus })[];
}

export interface GetFriendsResponse {
	users: User[];
}

export interface GetUserFriendsInputDto extends BasePaginationRequest {
	userId: number;
}

export interface AddFriendInputDto {
	userId: number;
	friendId: number;
}

export interface  AcceptFriendInputDto extends AddFriendInputDto  {
	isRejecting: boolean
};
