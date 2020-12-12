import { History } from 'history';
import { FC, useState } from 'react';

import { apiURL, createCookie } from '../utils';
import { User } from '../types';

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

export const Login: FC<loginProps> = ({ history }) => {
  const [state, setState] = useState({
    email: '',
    password: '',
    isSubmitting: false,
    message: '',
  });

  const { email, password, isSubmitting, message } = state;

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

      console.log(user);

      if (!success) {
        return setState({
          ...state,
          message: msg,
          isSubmitting: false,
        });
      }
      // expire in 30 minutes(same time as the cookie is invalidated on the backend)
      createCookie('token', token, 0.5);

      history.push({ pathname: '/session', state: user });
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

      <button disabled={isSubmitting} onClick={() => handleSubmit()}>
        {isSubmitting ? '.....' : 'login'}
      </button>
      <div className="message">{message}</div>
      <button
        onClick={() => {
          toRegister();
        }}
      >
        register
      </button>
    </div>
  );
};
