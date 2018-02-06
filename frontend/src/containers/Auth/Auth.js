import React, {Component} from 'react';
import {connect} from 'react-redux';
import axios from 'axios';
import FileSaver from 'file-saver';
import ReCAPTCHA from "react-google-recaptcha";
import * as actions from '../../store/actions/index';
import classes from './Auth.css';
import {API_URL} from "../../shared/const";
import Spinner from '../../components/UI/Spinner/Spinner';
import checkValidity from '../../shared/validation';
import Button from '../../components/UI/Button/Button';
import Input from '../../components/UI/Input/Input';
import Aux from '../../hoc/Aux/Aux';
import Modal from '../../components/UI/Modal/Modal';


class Auth extends Component {
    captcha;

    state = {
        register: false,
        loading: false,
        error: null,
        accData: null,
        regData: null,
        controls: {
            privateKey: {
                elementType: 'input',
                elementConfig: {
                    type: 'text',
                    placeholder: 'Private Key'
                },
                value: '',
                validation: {
                    required: true
                },
                valid: false,
                touched: false,
                errorMessage: null
            },
        }
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

    onSubmitHandler = (e) => {
        e.preventDefault();
        this.props.onAuth(this.state.controls.privateKey.value);
    };

    onSignUpHandler = () => {
        const conf = {headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        }};

        this.setState({regData: null, loading: true, error: null});

        axios.post(`${API_URL}/auth/sign-up`, {}, conf)
            .then(response => {
                this.setState({regData: response.data, loading: false, error: null});
                console.log(response)
            })
            .catch(error => {
                this.setState({regData: false, loading: false, error: error.response.data.message});
                console.log(error.response.data.message)
            })
    };

    saveAsDocHandler = () => {
        let text = '';
        Object.keys(this.state.regData).map(key => {return text += `${key} : ${this.state.regData[key]}\n`});
        const blob = new Blob([text], {type: "text/plain;charset=utf-8"});
        FileSaver.saveAs(blob, "my-account.txt");
    };

    modalCloseHandler = () => {
        this.setState({error: null});
    };

    render() {
        // form
        const formElementsArray = [];
        for ( let key in this.state.controls ) {
            formElementsArray.push( {
                id: key,
                config: this.state.controls[key]
            } );
        }

        let authForm = <form  className={classes.AuthForm} onSubmit={this.onSubmitHandler}>
            {
                formElementsArray.map(
                    formElement => <Input
                        errorMessage={formElement.config.errorMessage}
                        key={formElement.id}
                        elementType={formElement.config.elementType}
                        elementConfig={formElement.config.elementConfig}
                        value={formElement.config.value}
                        invalid={!formElement.config.valid}
                        shouldValidate={formElement.config.validation}
                        touched={formElement.config.touched}
                        changed={( event ) => this.inputChangedHandler( event, formElement.id )}
                    />
                )
            }
            <Button>Sign in</Button>
        </form>;
        // form /

        if (this.state.register) {
            if (!this.state.regData) {
                authForm = <form className={classes.AuthForm} onSubmit={(e) => {e.preventDefault(); this.captcha.execute(); this.onSignUpHandler()}}>
                    <ReCAPTCHA
                        ref="recaptcha"
                        size="invisible"
                        sitekey="6LfBAUQUAAAAAJi3xgNGMvz50-R5F2iTMgpW9J2q"
                        onChange={() =>{}}
                    />
                    <Button onClick={() => this.onSignUpHandler()}>Sign up</Button>
                </form>
            } else {
                authForm = <div className={classes.AuthForm}>
                    <h1 className={classes.Danger}>
                        Don't forget to save your account details.
                        <br/>
                        In case of loss, your account will be lost forever!
                    </h1>
                    {Object.keys(this.state.regData).map(key => <p key={key}>
                        {`${key} : ${this.state.regData[key]}`}
                    </p>)}
                    <Button onClick={() => this.saveAsDocHandler()}>
                        Save as doc
                    </Button>
                </div>
            }
        }

        let view = <div className={classes.Auth}>
            <label className={classes.preorderListTrigger}>
                <div>Sign in</div>
                <input type="checkbox"
                       checked={this.state.register}
                       onChange={() => {this.setState({register: !this.state.register})}}
                />
                <span />
                <div>Sign up</div>
            </label>

            {authForm}
        </div>;

        if (this.state.loading || this.props.loading) {
            view = <Spinner />
        }

        return <Aux>
                <Modal show={ this.state.error ? this.state.error : null }
                       modalClosed={() => this.modalCloseHandler()}>
                    {
                        this.state.error
                            ? <div className={classes.ModalContent}>
                                <h1>{this.state.error}</h1>
                                <Button onClick={() => this.modalCloseHandler()}>Ok</Button>
                            </div>
                            : null
                    }
                </Modal>
                <div  className={classes.Wrapper}>
                    {view}
                </div>
        </Aux>;
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.authKey,
        isAuth: state.auth.authKey !== null,
        loading: state.auth.loading
    }
};

const mapDispatchToProps = dispatch => {
    return {
        onAuth: (privateKey) => dispatch(actions.auth(privateKey)),
        // onLogout:() => dispatch(actions.logout())
    }
};

export default connect(mapStateToProps, mapDispatchToProps)(Auth);