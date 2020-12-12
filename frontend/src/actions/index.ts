import { Dispatch } from 'redux';

import { DeleteUserAction, UpdateUserAction } from './userActions';

export * from './userActions';
export * from './actionTypes';
export type Action = DeleteUserAction | UpdateUserAction;
export type ActionFunction = (dispatch: Dispatch) => void;
