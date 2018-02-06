import React, { Component } from 'react';

import Aux from '../Aux/Aux';
import classes from './Layout.css';
import Header from '../../components/Header/Header';

class Layout extends Component {
    render () {
        return (
            <Aux>
                <div className={classes.Layout}>
                    <Header />
                    <main>
                        {this.props.children}
                    </main>
                </div>
            </Aux>
        )
    }
}

export default Layout;