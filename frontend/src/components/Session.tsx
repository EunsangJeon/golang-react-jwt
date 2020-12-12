import { History } from 'history';
import { useState, useEffect, FC } from 'react';

import { User } from '../types';
import { apiURL, deleteCookie } from '../utils';

interface sessionProps {
  history: History;
}

interface getUserInfoResponse {
  success: boolean;
  user: User;
  msg?: string;
}

export const Session: FC<sessionProps> = ({ history }) => {
  const [state, setState] = useState({
    isFetching: false,
    message: '',
    user: {} as User,
  });

  const { isFetching, message, user } = state;

  const getUserInfo = async () => {
    setState({ ...state, isFetching: true, message: 'fetching details...' });
    try {
      const res: getUserInfoResponse = await fetch(`${apiURL}/session`, {
        method: 'GET',
        credentials: 'include',
        headers: {
          Accept: 'application/json',
          Authorization: document.cookie,
        },
      }).then((res) => res.json());

      const { success, user } = res;
      if (!success) {
        history.push('/login');
      }
      setState({ ...state, user, message: '', isFetching: false });
    } catch (e) {
      setState({ ...state, message: e.toString(), isFetching: false });
    }
  };
  const handleLogout = () => {
    deleteCookie('token');
    history.push('/login');
  };

  useEffect(() => {
    if (history.location.state) {
      return setState({ ...state, user: history.location.state as User });
    }
    getUserInfo();
  }, []);

  return (
    <div className="wrapper">
      <h1>Welcome, {user && user.name}</h1>
      {user && user.email}
      <div className="message">
        {isFetching ? 'fetching details..' : message}
      </div>

      <button
        style={{ height: '30px' }}
        onClick={() => {
          handleLogout();
        }}
      >
        logout
      </button>
    </div>
  );
};
