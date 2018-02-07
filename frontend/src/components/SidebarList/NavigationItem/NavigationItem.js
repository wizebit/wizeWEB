import React from 'react';
import { NavLink } from 'react-router-dom';
import classes from './NavigationItem.css';

const navigationItem = (props) => (
    <NavLink
        className={classes.NavigationItem}
        activeClassName={[classes.NavigationItem, classes.Active].join(' ')}
        exact
        to={props.link}
    >
        {props.children}
    </NavLink>
);

export default navigationItem;