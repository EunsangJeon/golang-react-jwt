export const createCookie = (
  cookieName: string,
  cookieValue: string,
  minuteToExpire: number
) => {
  const date = new Date();
  date.setTime(date.getTime() + minuteToExpire * 60 * 1000);
  document.cookie = `${cookieName} = ${cookieValue}; expires = ${date.toUTCString()}`;
};

export const checkCookie = (name: string) => {
  console.log(name);
  return document.cookie;
};

export const deleteCookie = (name: string) => {
  document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
};
