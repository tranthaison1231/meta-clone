export interface SignInRequest {
	email: string;
	password: string;
}

export interface SignUpRequest extends SignInRequest {
	gender: string;
	avatar?: string;
}
