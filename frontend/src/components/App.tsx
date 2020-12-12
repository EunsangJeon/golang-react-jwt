import { FC } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';

import '../styles/App.css';
import { Login, Register, Session } from '.';

export const App: FC = () => {
  return (
    <Router>
      <Route exact path="/" component={Login} />
      <Route path="/register" component={Register} />
      <Route path="/login" component={Login} />
      <Route path="/session" component={Session} />
    </Router>
  );
};

export default App;
