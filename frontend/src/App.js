import React, { Component } from 'react';
import { Route, Switch, withRouter, Redirect } from 'react-router-dom';
// import asyncComponent from './hoc/asyncComponent/asyncComponent';

// import Layout from './hoc/Layout/Layout';
import Auth from './containers/Auth/Auth';

class App extends Component {
  render() {
    return <div className="App">
        <Auth />
      </div>;
  }
}

export default App;
