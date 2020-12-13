import React, { FC, useCallback, useState } from 'react';
import { History } from 'history';
import { useSelector, useDispatch } from 'react-redux';
import Cookies from 'universal-cookie';

import { deleteUser, updateUser } from '../actions';
import { User } from '../types';
import { StoreState } from '../reducers';
import { apiURL } from '../utils';

interface sessionProps {
  history: History;
}

export const Session: FC<sessionProps> = (props: sessionProps) => {
  const [tokenInfo, setTokenInfo] = useState('');
  const [sessionInfo, setSessionInfo] = useState('');

  const dispatch = useDispatch();
  const cookies = new Cookies();

  const { history } = props;
  const user = useSelector<StoreState, User>((state) => state.user);

  const updateUserCallback = useCallback(
    (user: User) => dispatch(updateUser(user)),
    [dispatch]
  );

  const handleLogout = () => {
    cookies.remove('token');
    dispatch(deleteUser());
    history.push('/login');
  };

  const checkToken = () => {
    setTokenInfo(cookies.get('token'));
  };

  const getUserInfo = useCallback(async () => {
    try {
      const res = await fetch(`${apiURL}/session`, {
        method: 'GET',
        credentials: 'include',
        headers: {
          Accept: 'application/json',
          Authorization: cookies.getAll(),
        },
      }).then((res) => res.json());

      const { success, user, msg } = res;
      console.log(res);

      if (!success) {
        history.push('/login');
      }

      updateUserCallback(user);

      setSessionInfo(msg);
    } catch (e) {
      setSessionInfo(e.toString());
    }
  }, []);

  return (
    <div className="wrapper">
      <h1>Welcome, {user && user.name}</h1>
      {user && user.email}
      <button onClick={checkToken}>check token</button>
      <div className="tokenInfo">{tokenInfo}</div>
      <button onClick={getUserInfo}>check session to server</button>
      <div className="sessionInfo">{sessionInfo}</div>
      <button onClick={handleLogout}>logout</button>
    </div>
  );
};
