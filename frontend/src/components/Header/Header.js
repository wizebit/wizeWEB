import React from 'react';
import {Link} from 'react-router-dom';

import classes from './Header.css';

const header = props => <div className={classes.Header}>
    <div className={classes.HeaderContent}>
        <div className={classes.Logo}>
            <Link to="/">
                WizeBit
            </Link>
        </div>
        <div className={classes.Logout}>
            <Link to="/logout">
                Logout
            </Link>
        </div>
    </div>
</div>;

export default header;