import type { BasePaginationRequest } from './response';

export interface User {
	email: string;
	name: string;
	id: string;
	avatar: string;
	createdAt: string;
	updatedAt: string;
	firstName: string;
	lastName: string;
	gender: string;
	friends: User[] | null;
	friendRequests: FriendRequest[] | null;
}

export interface FriendRequest {
	id: string;
	userId: string;
	friendId: string;
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
	userId: string;
}

export interface AddFriendInputDto {
	userId: string;
	friendId: string;
}

export interface AcceptFriendInputDto extends AddFriendInputDto {
	isRejecting: boolean;
}
