import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Route, Switch, withRouter, Redirect } from 'react-router-dom';
import asyncComponent from './hoc/asyncComponent/asyncComponent';
import { authCheckState } from './store/actions/index';

import Auth from './containers/Auth/Auth';
import Layout from "./hoc/Layout/Layout";

const asyncLogout = asyncComponent(() => import('./containers/Auth/Logout/Logout'));
const asyncFilesList = asyncComponent(() => import('./containers/FilesList/FilesList'));
const asyncFileUpload = asyncComponent(() => import('./containers/FileUpload/FileUpload'));
const asyncWalletsList = asyncComponent(() => import('./containers/WalletsList/WalletsList'));
const asyncWalletCheck = asyncComponent(() => import('./containers/WalletCheck/WalletCheck'));

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
                  <Route exact path="/" component={asyncFilesList} />
                  <Route exact path="/logout" component={asyncLogout} />
                  <Route exact path="/upload-files" component={asyncFileUpload} />
                  <Route exact path="/wallets-list" component={asyncWalletsList} />
                  <Route exact path="/wallet-check" component={asyncWalletCheck} />
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