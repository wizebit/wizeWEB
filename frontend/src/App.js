import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Route, Switch, withRouter, Redirect } from 'react-router-dom';
import asyncComponent from './hoc/asyncComponent/asyncComponent';
import { authCheckState } from './store/actions/index';

import Auth from './containers/Auth/Auth';
import Layout from "./hoc/Layout/Layout";

const asyncPersonalArea = asyncComponent(() => import('./containers/PersonalArea/PersonalArea'));

class App extends Component {
    componentDidMount() {
        this.props.onTryAutoSignup();
    }

    render() {
      let routes = <Switch>
          <Route exact path="/" component={Auth} />
          <Redirect to="/" />
      </Switch>;

      if (this.props.isAuth) {
          routes = <Layout>
              <Switch>
                  <Route exact path="/" component={asyncPersonalArea} />
                  <Redirect to="/" />
              </Switch>
          </Layout>
      }

      return <div>
          {routes}
      </div>;
    }
}

// export default App;
const mapStateToProps = state => {
    return {
        isAuth: state.auth.authKey !== null
    }
};

const mapDispatchToProps = dispatch => {
    return {
        onTryAutoSignup: () => dispatch(authCheckState())
    }
};

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(App));