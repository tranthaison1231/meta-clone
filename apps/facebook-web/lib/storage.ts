export const getToken = () => {
  return localStorage.getItem('token');
};

export const setToken = (value: string) => {
  localStorage.setItem('token', value);
};

export const removeToken = () => {
  localStorage.removeItem('token');
};
