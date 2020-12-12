import React, { useState, FC } from 'react';
import { History } from 'history';

import { apiURL } from '../utils';

interface registerProps {
  history: History;
}

interface handleChangeEvent {
  target: {
    name: string;
    value: string;
  };
}

interface handleSubmitResponse {
  success: boolean;
  msg: string;
  errors: string[];
}

export const Register: FC<registerProps> = (props: registerProps) => {
  const { history } = props;

  const [state, setState] = useState({
    email: '',
    password: '',
    name: '',
    isSubmitting: false,
    message: '',
    errors: [''],
  });

  const { email, password, name, message, isSubmitting, errors } = state;

  const handleChange = async (event: handleChangeEvent) => {
    await setState({ ...state, [event.target.name]: event.target.value });
  };

  const handleSubmit = async () => {
    setState({ ...state, isSubmitting: true });

    const { email, password, name } = state;
    try {
      const res: handleSubmitResponse = await fetch(`${apiURL}/register`, {
        method: 'POST',
        body: JSON.stringify({
          email,
          password,
          name,
        }),
        headers: {
          'Content-Type': 'application/json',
        },
      }).then((res) => res.json());
      const { success, msg, errors } = res;

      if (!success) {
        return setState({
          ...state,
          message: msg,
          errors,
          isSubmitting: false,
        });
      }

      history.push('/login');
    } catch (error) {
      setState({ ...state, message: error.toString(), isSubmitting: false });
    }
  };

  return (
    <div className="wrapper">
      <h1>Register</h1>
      <input
        className="input"
        type="name"
        placeholder="Name"
        value={name}
        name="name"
        onChange={(event) => {
          handleChange(event);
        }}
      />
      <input
        className="input"
        type="text"
        placeholder="Email"
        value={email}
        name="email"
        onChange={(event) => {
          handleChange(event);
        }}
      />
      <input
        className="input"
        type="password"
        placeholder="Password"
        value={password}
        name="password"
        onChange={(event) => {
          handleChange(event);
        }}
      />

      <button disabled={isSubmitting} onClick={() => handleSubmit()}>
        {isSubmitting ? '...' : 'Sign Up'}
      </button>
      <div className="message">{message && <p>&bull; {message}</p>}</div>
      <div>
        {errors &&
          errors.map((error, id) => {
            return <p key={id}> &bull; {error}</p>;
          })}
      </div>
    </div>
  );
};
