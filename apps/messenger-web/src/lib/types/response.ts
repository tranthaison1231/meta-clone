export interface BaseResponseType<T> {
	code: number;
	data: T;
	status: string;
}

export interface BasePaginationRequest {
	page?: number;
	limit?: number;
	orderBy?: string;
}

export interface BasePaginationResponse<T> {
	items: T[];
	currentPage: number;
	pageSize: number;
	total: number;
}
