import { Dispatch } from 'redux';

import { ActionTypes } from './actionTypes';
import { User } from '../types';

export interface DeleteUserAction {
  type: ActionTypes.deleteUser;
}

export interface UpdateUserAction {
  type: ActionTypes.updateUser;
  payload: User;
}

export const deleteUser = () => {
  return (dispatch: Dispatch): void => {
    dispatch<DeleteUserAction>({
      type: ActionTypes.deleteUser,
    });
  };
};

export const updateUser = (user: User) => {
  return (dispatch: Dispatch): void => {
    dispatch<UpdateUserAction>({
      type: ActionTypes.updateUser,
      payload: user,
    });
  };
};
