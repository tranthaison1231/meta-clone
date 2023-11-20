export const validator = {
	email: /^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$/,
	password: /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$/,
	firstName: /\b([A-ZÀ-ÿ][-,a-z. ']+[ ]*)+/,
	lastName: /\b([A-ZÀ-ÿ][-,a-z. ']+[ ]*)+/,
	phone: /(84|0[3|5|7|8|9])+([0-9]{8})\b/g
};
