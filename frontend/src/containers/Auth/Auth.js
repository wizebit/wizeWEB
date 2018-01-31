import React, { Component } from 'react';
import { connect } from 'react-redux';

import Input from '../../components/UI/Input/Input';
import Button from '../../components/UI/Button/Button';
import Spinner from '../../components/UI/Spinner/Spinner';
import classes from './Auth.css';
import * as actions from '../../store/actions/index';
import checkValidity from '../../shared/validation';
import Aux from '../../hoc/Aux/Aux';
import Modal from '../../components/UI/Modal/Modal';

class Auth extends Component {
    state = {
        controls: {
            email: {
                elementType: 'input',
                elementConfig: {
                    type: 'email',
                    placeholder: 'Email'
                },
                value: '',
                validation: {
                    required: true,
                    isEmail: true
                },
                valid: false,
                touched: false,
                errorMessage: null
            },
            password: {
                elementType: 'input',
                elementConfig: {
                    type: 'password',
                    placeholder: 'Password'
                },
                value: '',
                validation: {
                    required: true
                },
                valid: false,
                touched: false,
                errorMessage: null
            },
        },
        isRegister: false
    };

    inputChangedHandler = (event, controlName) => {
        const updatedControls = {
            ...this.state.controls,
            [controlName]: {
                ...this.state.controls[controlName],
                value: event.target.value,
                valid: checkValidity(event.target.value, this.state.controls[controlName].validation, controlName).isValid,
                errorMessage: checkValidity(event.target.value, this.state.controls[controlName].validation, controlName).errorMessage,
                touched: true
            }
        };
        this.setState({controls: updatedControls});
    };

    submitHandler = (e) => {
        e.preventDefault();
        if (!this.state.isRegister) {
            this.props.onAuth(this.state.controls.userName.value, this.state.controls.password.value);
            console.log('login', this.state.controls.userName.value, this.state.controls.password.value);
        } else {
            console.log('register', this.state.controls.userName.value, this.state.controls.password.value);

        }
    };

    changeFormHandler = () => {
        this.setState({isRegister: !this.state.isRegister});
    };

    modalCloseHandler = () => {
        this.props.onLogout('/cart');
    };

    render () {
        // form
        const formElementsArray = [];
        for ( let key in this.state.controls ) {
            formElementsArray.push( {
                id: key,
                config: this.state.controls[key]
            } );
        }

        let form = formElementsArray.map( formElement => (
            <Input
                errorMessage={formElement.config.errorMessage}
                key={formElement.id}
                elementType={formElement.config.elementType}
                elementConfig={formElement.config.elementConfig}
                value={formElement.config.value}
                invalid={!formElement.config.valid}
                shouldValidate={formElement.config.validation}
                touched={formElement.config.touched}
                changed={( event ) => this.inputChangedHandler( event, formElement.id )} />
        ) );
        // from /

        if (this.props.loading) {
            form = <Spinner />;
        }

        //On form change
        let nowText = "Sign In",
            altText = "Sign Up";
        if (this.state.isRegister) {
            nowText = "Sign Up";
            altText = "Sign In";
        }
        //On form change /


        return (

            <Aux>
                <Modal
                    show={ this.props.error ? this.props.error : null }
                    modalClosed={() => this.modalCloseHandler}>
                    {
                        this.props.error
                        ? this.props.error
                        : null
                    }
                </Modal>
                <div className={classes.Auth}>
                    <form className={classes.AuthForm} onSubmit={this.submitHandler}>
                        <h1>{nowText}</h1>

                        {form}

                        <div className={classes.ButtonSection}>
                            <Button>
                                {nowText}
                            </Button>
                            <Button
                                onClick={(e) => {
                                    e.preventDefault();
                                    this.changeFormHandler()
                                }}
                            >
                                {altText}
                            </Button>
                        </div>
                    </form>
                </div>
            </Aux>
        );
    }
}

const mapStateToProps = state => {
    return {
        isAuth: state.auth.authKey !== null,
        loading: state.auth.loading,
        error: state.auth.error,
    }
};

const mapDispatchToProps = dispatch => {
    return {
        onAuth: (username, password) => dispatch(actions.auth(username, password)),
        onLogout:() => dispatch(actions.logout())
    }
};

export default connect(mapStateToProps, mapDispatchToProps)(Auth);