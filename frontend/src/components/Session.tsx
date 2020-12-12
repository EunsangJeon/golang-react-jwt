import { History } from 'history';
import { FC, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { Dispatch } from 'redux';

import { deleteUser } from '../actions';
import { User } from '../types';
import { StoreState } from '../reducers';
import { checkCookie, deleteCookie } from '../utils';

interface sessionProps {
  history: History;
}

export const Session: FC<sessionProps> = ({ history }) => {
  const user = useSelector<StoreState, User>((state) => state.user);
  if (user.name === '') {
    history.push('/login');
  }

  const [tokenInfo, setTokenInfo] = useState('');

  const dispatch: Dispatch<any> = useDispatch();

  const handleLogout = () => {
    deleteCookie('token');
    dispatch(deleteUser());
    history.push('/login');
  };

  const checkToken = () => {
    setTokenInfo(checkCookie('token'));
  };

  return (
    <div className="wrapper">
      <h1>Welcome, {user && user.name}</h1>
      {user && user.email}
      <button onClick={checkToken}>check token</button>
      <div className="tokenInfo">{tokenInfo}</div>
      <button style={{ height: '30px' }} onClick={handleLogout}>
        logout
      </button>
    </div>
  );
};
