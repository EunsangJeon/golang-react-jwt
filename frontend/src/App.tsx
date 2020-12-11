import { BrowserRouter as Router, Route } from 'react-router-dom';

import './App.css';
import { Login } from './components/Login';
import { Register } from './components/Register';
import { Session } from './components/Session';

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
