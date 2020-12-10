export const createCookie = (
  cookieName: string,
  cookieValue: string,
  hourToExpire: number
) => {
  const date = new Date();
  date.setTime(date.getTime() + hourToExpire * 60 * 60 * 1000);
  document.cookie = `${cookieName} = ${cookieValue}; expires = ${date.toUTCString()}`;
};

export const deleteCookie = (name: string) => {
  document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
};
