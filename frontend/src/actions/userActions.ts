import { Dispatch } from 'redux';

import { ActionTypes } from './actionTypes';
import { User } from '../types';

export interface FetchUserAction {
  type: ActionTypes.fetchUser;
}

export interface UpdateUserAction {
  type: ActionTypes.updateUser;
  payload: User;
}

export const fetchUser = () => {
  return (dispatch: Dispatch) => {
    dispatch<FetchUserAction>({
      type: ActionTypes.fetchUser,
    });
  };
};

export const updateUser = (user: User) => {
  return (dispatch: Dispatch) => {
    dispatch<UpdateUserAction>({
      type: ActionTypes.updateUser,
      payload: user,
    });
  };
};
