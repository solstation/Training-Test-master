import Home from './components/home'
import Favorites from './components/favorites'

import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

// Here, we will have the router of the app
function App() {
  return (
    <div className="App">
      <Router>
        <header className="shadow">
          <nav>
            <ul>
              <li>
                <Link to='/'>Home</Link>
              </li>
              <li>
                <Link to='/favorites'>Favorites</Link>
              </li>
            </ul>
          </nav>
        </header>
        <Switch>
          <Route path='/favorites'>
            <Favorites />
          </Route>
          <Route path='/'>
            <Home />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
