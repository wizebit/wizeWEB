import React from 'react';

import classes from './Modal.css';
import Aux from '../../../hoc/Aux/Aux';
import Backdrop from '../Backdrop/Backdrop';

const modal = (props) => {
        return (
            <Aux>
                <Backdrop show={props.show} clicked={props.modalClosed} />
                <div className={classes.Modal}>
                    <span className={classes.XMark} onClick={props.modalClosed}>x</span>
                    {props.children}
                </div>
            </Aux>
        )
};

export default modal;