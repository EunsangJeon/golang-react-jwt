import { BrowserRouter as Router, Route } from 'react-router-dom';

import './App.css';
import { Login, Register, Session } from './components';

function App() {
  return (
    <Router>
      <Route exact path="/" component={Login} />
      <Route path="/register" component={Register} />
      <Route path="/login" component={Login} />
      <Route path="/session" component={Session} />
    </Router>
  );
}

export default App;
