import React, { Component } from 'react';
import { connect } from 'react-redux';

import Aux from '../Aux/Aux';
import classes from './Layout.css';
import Toolbar from '../../components/Navigation/Toolbar/Toolbar';
import Spinner from '../../components/UI/Spinner/Spinner';

class Layout extends Component {
    render () {
        const toolbar = this.props.onlineOrdering
            ? <Toolbar
                isAuth={this.props.isAuth}
                cartQty={this.props.cartQty}
                locReady={this.props.locReady}
                delivery={this.props.delivery}
                profileName={this.props.profileName}
                profileImage={this.props.profileImage} />
            : null;

        let content = <div className={classes.SpinnerToMiddle}>
            <Spinner />
        </div>;

        if (this.props.locReady) {
            content = <div>
                {toolbar}

                <main className={classes.Content}>
                    {this.props.children}
                </main>
            </div>;
        }

        return (
            <Aux>
                <div className={classes.Layout}>
                    {content}
                </div>
            </Aux>
        )
    }
}

const mapStateToProps = state => {
    return {
        isAuth: state.auth.authKey !== null,
        cartQty: state.cart.qty,
        locReady: !state.location.loading,
        delivery: state.location.deliveryEnable,
        onlineOrdering: state.location.onlineOrdering,
        profileName: state.auth.title,
        profileImage: state.auth.profileImage
    }
};

export default connect(mapStateToProps)(Layout);