export function constructPayload<T>(action: string, payload: T) {
	return JSON.stringify({
		action,
		payload
	});
}
