export const createCookie = (
  cookieName: string,
  cookieValue: string,
  minuteToExpire: number
): void => {
  const date = new Date();
  date.setTime(date.getTime() + minuteToExpire * 60 * 1000);
  document.cookie = `${cookieName} = ${cookieValue}; expires = ${date.toUTCString()}`;
};

export const checkCookie = (name: string): string => {
  console.log(name);
  return document.cookie;
};

export const deleteCookie = (name: string): void => {
  document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
};
