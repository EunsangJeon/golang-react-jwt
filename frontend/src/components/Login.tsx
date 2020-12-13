import { History } from 'history';
import React, { FC, useState, useCallback } from 'react';
import { useDispatch } from 'react-redux';
import Cookies from 'universal-cookie';

import { updateUser } from '../actions';
import { User } from '../types';
import { apiURL, SECOND } from '../utils';

interface loginProps {
  history: History;
}

interface handleChangeEvent {
  target: {
    name: string;
    value: string;
  };
}

interface loginResponse {
  token: string;
  success: boolean;
  msg: string;
  user: User;
}

export const Login: FC<loginProps> = (props: loginProps) => {
  const [state, setState] = useState({
    email: '',
    password: '',
    isSubmitting: false,
    message: '',
  });
  const { history } = props;
  const { email, password, isSubmitting, message } = state;
  const dispatch = useDispatch();
  const cookies = new Cookies();

  const updateUserCallback = useCallback(
    (user: User) => dispatch(updateUser(user)),
    [dispatch]
  );

  const handleChange = async (event: handleChangeEvent) => {
    const { name, value } = event.target;
    await setState({ ...state, [name]: value });
  };

  const toRegister = () => {
    history.push({ pathname: '/register' });
  };

  const handleSubmit = async () => {
    setState({ ...state, isSubmitting: true });

    const { email, password } = state;
    try {
      const res: loginResponse = await fetch(`${apiURL}/login`, {
        method: 'POST',
        body: JSON.stringify({
          email,
          password,
        }),
        headers: {
          'Content-Type': 'application/json',
        },
      }).then((res) => res.json());

      const { token, success, msg, user } = res;

      if (!success) {
        return setState({
          ...state,
          message: msg,
          isSubmitting: false,
        });
      }
      // expire in 10 second(same time as the cookie is invalidated on the backend)
      cookies.set('token', token, {
        path: '/',
        expires: new Date(Date.now() + 10 * SECOND),
      });

      updateUserCallback(user);

      history.push({ pathname: '/session' });
    } catch (error) {
      setState({ ...state, message: error.toString(), isSubmitting: false });
    }
  };

  return (
    <div className="wrapper">
      <h1>Login</h1>
      <input
        className="input"
        type="text"
        placeholder="email"
        value={email}
        name="email"
        onChange={(event) => {
          handleChange(event);
        }}
      />

      <input
        className="input"
        type="password"
        placeholder="password"
        value={password}
        name="password"
        onChange={(event) => {
          handleChange(event);
        }}
      />

      <button disabled={isSubmitting} onClick={handleSubmit}>
        {isSubmitting ? '...' : 'login'}
      </button>
      <div className="message">{message}</div>
      <button onClick={toRegister}>register</button>
    </div>
  );
};
