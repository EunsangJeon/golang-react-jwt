import { combineReducers } from 'redux';

import { userReducer } from './userReducer';
import { User } from '../types';

export interface StoreState {
  user: User;
}

export const reducers = combineReducers<StoreState>({
  user: userReducer,
});
