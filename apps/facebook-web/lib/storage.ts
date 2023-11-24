export const getToken = () => {
  const token = useCookie('token');
  return token.value;
};

export const setToken = (value: string) => {
  const token = useCookie('token');
  token.value = value;
};

export const removeToken = () => {
  const token = useCookie('token');
  token.value = '';
};
