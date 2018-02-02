import React, { Component } from 'react';
// import { connect } from 'react-redux';
import { Route, Switch, withRouter, Redirect } from 'react-router-dom';
// import asyncComponent from './hoc/asyncComponent/asyncComponent';

import Auth from './containers/Auth/Auth';

class App extends Component {
  render() {
      let routes = <Switch>
          <Route exact path="/" component={Auth} />
          <Redirect to="/" />
      </Switch>;

      return <div>
          {routes}
      </div>;
  }
}

export default App;
