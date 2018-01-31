import React, {Component} from 'react';

import classes from './Input.css';
import Aux from '../../../hoc/Aux/Aux';
import Backdrop from '../../UI/Backdrop/Backdrop';

class Input extends Component {
    state = {
        datepickerShow: false,
        errorMessage: null,
        uploadImage: null
    };

    showDatepicker = () => {
        this.setState({datepickerShow: !this.state.datepickerShow});
    };

    render() {
        let inputElement = null;
        const inputClasses = [classes.InputElement];

        if (this.props.invalid && this.props.shouldValidate && this.props.touched) {
            inputClasses.push(classes.Invalid);
        }

        if (!this.props.extraOptions) {
            switch (this.props.elementType) {
                case ( 'input' ):
                    inputElement = (
                        <div>
                            <input
                                className={inputClasses.join(' ')}
                                {...this.props.elementConfig}
                                value={this.props.value}
                                onChange={this.props.changed}/>
                            <div className={classes.ErrorMessage}>
                                {this.props.invalid ? this.props.errorMessage : null}
                            </div>
                        </div>
                    );
                    break;
                case ( 'textarea' ):
                    inputElement = (
                        <div>
                            <textarea
                                className={inputClasses.join(' ')}
                                {...this.props.elementConfig}
                                value={this.props.value}
                                onChange={this.props.changed} />

                        </div>
                    );
                    break;
                case ( 'select' ):
                    inputElement = (
                        <div>
                            <select
                                className={inputClasses.join(' ')}
                                value={this.props.value}
                                onChange={this.props.changed}>
                                {this.props.elementConfig.options.map(option => (
                                    <option key={option.value} value={option.value}>
                                        {option.displayValue}
                                    </option>
                                ))}
                            </select>
                            <div className={classes.ErrorMessage}>
                                {this.props.invalid ? this.props.errorMessage : null}
                            </div>
                        </div>
                    );
                    break;
                default:
                    inputElement = (
                        <div>
                            <input
                                className={inputClasses.join(' ')}
                                {...this.props.elementConfig}
                                value={this.props.value}
                                onChange={this.props.changed}/>
                            <div className={classes.ErrorMessage}>
                                {this.props.invalid ? this.props.errorMessage : null}
                            </div>
                        </div>
                    );
                    break;
            }
        }

        return (
            <Aux>
                <Backdrop transparent show={this.state.datepickerShow} clicked={this.showDatepicker}/>
                <div className={classes.Input}>
                    <label className={classes.Label}>{this.props.label}</label>
                    {inputElement}
                </div>
            </Aux>
        );
    }
}

export default Input;