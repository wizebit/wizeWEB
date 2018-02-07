import React, { Component } from 'react';

import Aux from '../Aux/Aux';
import classes from './Layout.css';
import Header from '../../components/Header/Header';
import SidebarList from '../../components/SidebarList/SidebarList';

class Layout extends Component {
    render () {
        return (
            <Aux>
                <div className={classes.Layout}>
                    <Header />
                    <main>
                        <SidebarList />

                        <article>
                            {this.props.children}
                        </article>
                    </main>
                </div>
            </Aux>
        )
    }
}

export default Layout;