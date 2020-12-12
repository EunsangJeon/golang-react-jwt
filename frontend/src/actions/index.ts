import { DeleteUserAction, UpdateUserAction } from './userActions';

export * from './actionTypes';
export type Action = DeleteUserAction | UpdateUserAction;
