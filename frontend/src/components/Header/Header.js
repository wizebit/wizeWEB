import React from 'react';

import classes from './Header.css';

const header = props => <div className={classes.Header}>
    <div className={classes.HeaderContent}>
        <div className={classes.Logo}>
            WizeBit
        </div>
    </div>
</div>;

export default header;