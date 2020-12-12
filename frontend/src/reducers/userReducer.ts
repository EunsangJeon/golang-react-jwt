import { Action, ActionTypes } from '../actions';
import { User } from '../types';

const initUserState: User = {
  email: '',
  name: '',
  created_at: '',
  updated_at: '',
};

export const userReducer = (
  state: User = initUserState,
  action: Action
): User => {
  switch (action.type) {
    case ActionTypes.updateUser:
      return action.payload;
    case ActionTypes.deleteUser:
      return initUserState;
    default:
      return state;
  }
};
