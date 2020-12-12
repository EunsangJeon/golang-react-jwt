import { createStore, applyMiddleware, Store } from 'redux';
import { persistStore, persistReducer, Persistor } from 'redux-persist';
import storage from 'redux-persist/lib/storage';
import thunk from 'redux-thunk';

import { reducers } from '../reducers';

const persistConfig = {
  key: 'root',
  storage,
};

const persistedReducer = persistReducer(persistConfig, reducers);

export const configureStore = (): { persistor: Persistor; store: Store } => {
  const store = createStore(persistedReducer, applyMiddleware(thunk));
  const persistor = persistStore(store);

  return { persistor, store };
};
