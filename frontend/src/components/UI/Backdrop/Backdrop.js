import React from 'react';

import classes from './Backdrop.css';

const backdrop = (props) => (
    props.show ? <div className={ props.transparent ? [classes.Transparent, classes.Backdrop].join(' ') : classes.Backdrop }
        onClick={props.clicked}></div> : null
);

export default backdrop;